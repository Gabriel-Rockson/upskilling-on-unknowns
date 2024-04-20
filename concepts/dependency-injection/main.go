package main

// Followed on https://medium.com/@LittleLeverages/what-even-is-dependency-injection-using-go-7f85724bbdb2

type Entitlement struct {
  AccountId   string
  CanTransfer bool
}

type EntitlementsClient interface {
  GetEntitlement(accountId string) (Entitlement, error)
}

type Account struct {
  UserId     string
  AccountId  string
}

type AccountsRepository interface {
  GetByUserId(userId string) ([]Account, error)
}

type Service interface {
  GetTransferAccounts(userId string) ([]Account, error)
}

type service struct {
  accountsRepository AccountsRepository
  entitlementClient EntitlementsClient
}

// DEPENDENCY INJECTION
func NewService(accountsRepository AccountsRepository, entitlementsClient EntitlementsClient) service {
  return service{
    accountsRepository,
    entitlementsClient,
  }
}

func (s *service) GetTransferAccounts(userId string) ([]Account, error) {
  accounts, err := s.accountsRepository.GetByUserId(userId)
  if err != nil {
    return nil, err
  }

  output := []Account{}

  for _, a := range accounts {
    entitlement, err := s.entitlementClient.GetEntitlement(a.AccountId)
    if err != nil {
      return nil, err
    }

    if entitlement.CanTransfer {
      output = append(output, a)
    }
  }

  return output, nil
}

