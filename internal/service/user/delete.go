package user

import "context"

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.userRepository.Delete(ctx, id)
	return err
}
