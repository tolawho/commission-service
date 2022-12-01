package services

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	pnt_transaction "medici.vn/commission-serivce/enums/pnt-transaction"
	"medici.vn/commission-serivce/models"
	"medici.vn/commission-serivce/repository"
)

// PntDailyCommissionService  is contract what pntDailyCommissionService can do to db
type PntDailyCommissionService interface {
	Calculator(id uint) models.PntContract
}

type pntDailyCommissionService struct {
	pntDailyCommissionRepository   repository.PntDailyCommissionRepository
	pntContractRepository          repository.PntContractRepository
	pntCommissionFormulaRepository repository.PntCommissionFormulaRepository
	pntContractProductRepository   repository.PntContractProductRepository
	pntPolicyRepository            repository.PntPolicyRepository
	pntAgencyTreeRepository        repository.PntAgencyTreeRepository
	agencyRepository               repository.AgencyRepository
	pntTransactionRepository       repository.PntTransactionRepository
	connection                     *gorm.DB
}

func (p pntDailyCommissionService) Calculator(id uint) models.PntContract {
	var pntContract = p.pntContractRepository.FindById(id)
	if pntContract.ID == 0 {
		return pntContract
	}

	var pntContractProducts = p.pntContractProductRepository.FindByContractId(pntContract.ID)
	var agency = p.agencyRepository.FindById(pntContract.AgencyId)
	var policy = p.pntPolicyRepository.FindActive()
	var levels = []string{"CBDO", "CEO", "CHAIRMAN"}

	p.processCalculator(pntContract, pntContractProducts, agency, nil, policy)

	for agency.ID != 0 {
		var agencyTree = p.pntAgencyTreeRepository.FindByAgencyId(agency.ID)
		if agencyTree != nil {
			var agencyChild = agency
			agency = p.agencyRepository.FindById(agencyTree.ParentId)
			if agency.ID <= 2 || funk.Contains(levels, agency.PntAgencyLevelCode) {
				break
			}
			p.processCalculator(pntContract, pntContractProducts, agency, agencyChild, policy)
		}
	}

	return pntContract
}

func (p pntDailyCommissionService) processCalculator(
	pntContract models.PntContract,
	pntContractProducts []*models.PntContractProduct,
	agency *models.Agency,
	agencyChild *models.Agency,
	policy models.PntPolicy,
) {
	var commission float32 = 0
	for _, pntContractProduct := range pntContractProducts {
		var value = pntContractProduct.CommissionRate
		var formula *models.PntCommissionFormula
		var beforeFormula *models.PntCommissionFormula
		if policy.Status == "ON" {
			formula = p.pntCommissionFormulaRepository.FindFormula(
				models.PntCommissionFormula{
					LevelCode:    p.FindLevel(agency),
					PntProductId: pntContractProduct.PntProductId,
					PntPolicyId:  policy.ID,
				},
			)
			if agencyChild != nil && agencyChild.ID != 0 {
				beforeFormula = p.pntCommissionFormulaRepository.FindFormula(
					models.PntCommissionFormula{
						LevelCode:    p.FindLevel(agencyChild),
						PntProductId: pntContractProduct.PntProductId,
						PntPolicyId:  policy.ID,
					},
				)
				if formula == nil || formula.ID == 0 {
					return
				}
			}
			if formula != nil && formula.ID != 0 {
				value = formula.Value
				if beforeFormula != nil && beforeFormula.ID != 0 {
					value = formula.Value - beforeFormula.Value
				}
			}
		}
		commission += (pntContractProduct.Amount - pntContractProduct.Tax) * value / 100
	}
	var sourceId = agency.ID
	var amount float32 = 0
	var sysCommission float32 = 0

	if agencyChild != nil && agencyChild.ID != 0 {
		sourceId = agencyChild.ID
		sysCommission = commission
	} else {
		amount = commission
	}

	tx := p.connection.Begin()
	_, err1 := p.pntDailyCommissionRepository.FirstOrCreate(
		models.PntDailyCommission{
			AgencyId:      agency.ID,
			PntContractId: pntContract.ID,
		},
		models.PntDailyCommission{
			Day:             pntContract.CreatedAt,
			SourceId:        sourceId,
			Amount:          amount,
			SysCommission:   sysCommission,
			LevelCode:       p.FindLevel(agency),
			SourceLevelCode: p.FindLevel(agencyChild),
			SourceModel:     "Agency",
			IsOldData:       false,
			PolicyId:        policy.ID,
		})

	_, err2 := p.pntTransactionRepository.FirstOrCreate(
		models.PntTransaction{
			PntContractId: pntContract.ID,
			AgencyId:      agency.ID,
		},
		models.PntTransaction{
			Note:          fmt.Sprintf("Ghi nhận hoa hồng cho Agency %d từ hợp đồng %d", agency.ID, pntContract.ID),
			AgencyId:      agency.ID,
			PntContractId: pntContract.ID,
			Type:          pnt_transaction.TYPE_COMMISSION,
			Status:        pnt_transaction.STATUS_SUCCESSFUL,
			Amount:        commission,
		})
	if err1 != nil || err2 != nil {
		spew.Dump(err1, "aaaa")
		spew.Dump(err2, "bbb")
		tx.Rollback()
		return
	}
	tx.Commit()
}

func (p pntDailyCommissionService) FindLevel(agency *models.Agency) string {
	if agency == nil || agency.ID == 0 {
		return ""
	}
	var level = agency.PntAgencyLevelCode
	if level == "" {
		level = agency.PntFullTimeLevelCode
	}
	return level
}

// NewPntDailyCommissionService is creates a new instance of PntDailyCommissionService
func NewPntDailyCommissionService(
	pntDailyCommissionRepo repository.PntDailyCommissionRepository,
	pntContractRepo repository.PntContractRepository,
	pntCommissionFormulaRepo repository.PntCommissionFormulaRepository,
	pntContractProductRepo repository.PntContractProductRepository,
	pntAgencyTreeRepo repository.PntAgencyTreeRepository,
	pntPolicyRepo repository.PntPolicyRepository,
	agencyRepo repository.AgencyRepository,
	pntTransactionRepo repository.PntTransactionRepository,
	db *gorm.DB,
) PntDailyCommissionService {
	return &pntDailyCommissionService{
		pntDailyCommissionRepository:   pntDailyCommissionRepo,
		pntContractRepository:          pntContractRepo,
		pntCommissionFormulaRepository: pntCommissionFormulaRepo,
		pntContractProductRepository:   pntContractProductRepo,
		pntAgencyTreeRepository:        pntAgencyTreeRepo,
		pntPolicyRepository:            pntPolicyRepo,
		agencyRepository:               agencyRepo,
		pntTransactionRepository:       pntTransactionRepo,
		connection:                     db,
	}
}
