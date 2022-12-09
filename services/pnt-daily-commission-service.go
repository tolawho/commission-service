package services

import (
	"fmt"
	"github.com/thoas/go-funk"
	pntLevelPartTime "medici.vn/commission-serivce/enums/pnt-level-part-time"
	pntTransaction "medici.vn/commission-serivce/enums/pnt-transaction"
	"medici.vn/commission-serivce/models"
	"medici.vn/commission-serivce/repository"
)

// PntDailyCommissionService  is contract what pntDailyCommissionService can do to db
type PntDailyCommissionService interface {
	Calculator(id uint) (models.PntContract, error)
}

type pntDailyCommissionService struct {
	pntDailyCommissionRepository   repository.PntDailyCommissionRepository
	pntContractRepository          repository.PntContractRepository
	pntCommissionFormulaRepository repository.PntCommissionFormulaRepository
	pntPolicyRepository            repository.PntPolicyRepository
	pntAgencyTreeRepository        repository.PntAgencyTreeRepository
	agencyRepository               repository.AgencyRepository
	pntTransactionRepository       repository.PntTransactionRepository
}

func (p pntDailyCommissionService) Calculator(id uint) (models.PntContract, error) {
	var pntContract, err = p.pntContractRepository.First(models.PntContract{ID: id})

	if err != nil {
		return pntContract, err
	}

	var pntContractProducts = pntContract.PntContractProducts
	var agency = pntContract.Agency
	var policy = p.pntPolicyRepository.FindActive()

	if len(pntContractProducts) == 0 || agency.ID == 0 || policy.ID == 0 {
		return pntContract, nil
	}

	err = p.processCalculator(pntContract, pntContractProducts, agency, nil, policy)
	if err != nil {
		return pntContract, err
	}

	// top sales
	var levels = []string{pntLevelPartTime.CHAIRMAN}
	for agency.ID != 0 {
		var agencyTree = p.pntAgencyTreeRepository.FindByAgencyId(agency.ID)
		if agencyTree != nil {
			var agencyChild = agency
			agency = p.agencyRepository.FindById(agencyTree.ParentId)
			if agency.ID <= 5 || funk.Contains(levels, agency.PntLvPartTime) {
				break
			}
			err = p.processCalculator(pntContract, pntContractProducts, agency, agencyChild, policy)
			if err != nil {
				return pntContract, err
			}
		}
	}

	return pntContract, nil
}

func (p pntDailyCommissionService) processCalculator(
	pntContract models.PntContract,
	pntContractProducts []*models.PntContractProduct,
	agency *models.Agency,
	agencyChild *models.Agency,
	policy models.PntPolicy,
) error {
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
					return nil
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

	if _, err := p.pntDailyCommissionRepository.FirstOrCreate(
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
		}); err != nil {
		return err
	}

	if _, err := p.pntTransactionRepository.FirstOrCreate(
		models.PntTransaction{
			PntContractId: pntContract.ID,
			AgencyId:      agency.ID,
			Type:          pntTransaction.TYPE_COMMISSION,
		},
		models.PntTransaction{
			Note:          fmt.Sprintf("Ghi nhận hoa hồng cho Agency %s từ hợp đồng %d", agency.Code, pntContract.ID),
			AgencyId:      agency.ID,
			PntContractId: pntContract.ID,
			Type:          pntTransaction.TYPE_COMMISSION,
			Status:        pntTransaction.STATUS_SUCCESSFUL,
			Amount:        commission,
		}); err != nil {
		return err
	}

	return nil
}

func (p pntDailyCommissionService) FindLevel(agency *models.Agency) string {
	if agency == nil || agency.ID == 0 {
		return ""
	}
	var level = agency.PntLvPartTime
	if level == "" {
		level = agency.PntLvFullTime
	}
	return level
}

// NewPntDailyCommissionService is creates a new instance of PntDailyCommissionService
func NewPntDailyCommissionService(
	pntDailyCommissionRepo repository.PntDailyCommissionRepository,
	pntContractRepo repository.PntContractRepository,
	pntCommissionFormulaRepo repository.PntCommissionFormulaRepository,
	pntAgencyTreeRepo repository.PntAgencyTreeRepository,
	pntPolicyRepo repository.PntPolicyRepository,
	agencyRepo repository.AgencyRepository,
	pntTransactionRepo repository.PntTransactionRepository,
) PntDailyCommissionService {
	return &pntDailyCommissionService{
		pntDailyCommissionRepository:   pntDailyCommissionRepo,
		pntContractRepository:          pntContractRepo,
		pntCommissionFormulaRepository: pntCommissionFormulaRepo,
		pntAgencyTreeRepository:        pntAgencyTreeRepo,
		pntPolicyRepository:            pntPolicyRepo,
		agencyRepository:               agencyRepo,
		pntTransactionRepository:       pntTransactionRepo,
	}
}
