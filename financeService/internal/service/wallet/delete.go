package wallet

import "context"

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.walletRepository.Delete(ctx, id)
	return err
}
