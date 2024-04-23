package apiserver

import "github.com/esivanov203/go-rest-api/internal/store"

func (s *ApiServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}
