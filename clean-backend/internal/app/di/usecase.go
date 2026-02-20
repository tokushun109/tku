package di

import (
	clockInfra "github.com/tokushun109/tku/clean-backend/internal/infra/clock"
	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
	cryptoInfra "github.com/tokushun109/tku/clean-backend/internal/infra/crypto"
	mailInfra "github.com/tokushun109/tku/clean-backend/internal/infra/mail/sendgrid"
	uuidInfra "github.com/tokushun109/tku/clean-backend/internal/infra/uuid"
	usecase "github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseCategory "github.com/tokushun109/tku/clean-backend/internal/usecase/category"
	usecaseContact "github.com/tokushun109/tku/clean-backend/internal/usecase/contact"
	usecaseHealth "github.com/tokushun109/tku/clean-backend/internal/usecase/health"
	usecaseSalesSite "github.com/tokushun109/tku/clean-backend/internal/usecase/sales_site"
	usecaseSession "github.com/tokushun109/tku/clean-backend/internal/usecase/session"
	usecaseSkillMarket "github.com/tokushun109/tku/clean-backend/internal/usecase/skill_market"
	usecaseSns "github.com/tokushun109/tku/clean-backend/internal/usecase/sns"
	usecaseTag "github.com/tokushun109/tku/clean-backend/internal/usecase/tag"
	usecaseTarget "github.com/tokushun109/tku/clean-backend/internal/usecase/target"
	usecaseUser "github.com/tokushun109/tku/clean-backend/internal/usecase/user"
)

type usecases struct {
	health      usecaseHealth.Usecase
	category    usecaseCategory.Usecase
	target      usecaseTarget.Usecase
	tag         usecaseTag.Usecase
	sns         usecaseSns.Usecase
	salesSite   usecaseSalesSite.Usecase
	skillMarket usecaseSkillMarket.Usecase
	contact     usecaseContact.Usecase
	session     usecaseSession.Usecase
	user        usecaseUser.Usecase
}

func newUsecases(repos *repositories, cfg *config.Config, txManager usecase.TxManager) (*usecases, error) {
	// 入力側の依存関係のチェック
	if err := requireStructFieldsNonNil("repositories", repos); err != nil {
		return nil, err
	}
	if err := requireNonNil("config", cfg); err != nil {
		return nil, err
	}
	if err := requireNonNil("txManager", txManager); err != nil {
		return nil, err
	}

	uuidGen := uuidInfra.NewGenerator()
	if err := requireNonNil("uuidGenerator", uuidGen); err != nil {
		return nil, err
	}
	clock := clockInfra.NewClock()
	if err := requireNonNil("clock", clock); err != nil {
		return nil, err
	}
	passwordHasher := cryptoInfra.NewPasswordHasherSHA1()
	if err := requireNonNil("passwordHasher", passwordHasher); err != nil {
		return nil, err
	}
	contactNotifier := mailInfra.NewContactNotifier(cfg.Env, cfg.SendGridAPIKey, repos.user)
	if err := requireNonNil("contactNotifier", contactNotifier); err != nil {
		return nil, err
	}
	sessionUC := usecaseSession.New(repos.session, cfg.SessionTTL, clock)
	if err := requireNonNil("sessionUsecase", sessionUC); err != nil {
		return nil, err
	}

	ucs := &usecases{
		health:      usecaseHealth.New(repos.health),
		category:    usecaseCategory.New(repos.category, uuidGen),
		target:      usecaseTarget.New(repos.target, uuidGen),
		tag:         usecaseTag.New(repos.tag, uuidGen),
		sns:         usecaseSns.New(repos.sns, uuidGen),
		salesSite:   usecaseSalesSite.New(repos.salesSite, uuidGen),
		skillMarket: usecaseSkillMarket.New(repos.skillMarket, uuidGen),
		contact:     usecaseContact.New(repos.contact, contactNotifier),
		session:     sessionUC,
		user:        usecaseUser.New(repos.user, repos.session, sessionUC, passwordHasher, uuidGen, clock, txManager),
	}

	// 出力側の依存関係のチェック
	if err := requireStructFieldsNonNil("usecases", ucs); err != nil {
		return nil, err
	}

	return ucs, nil
}
