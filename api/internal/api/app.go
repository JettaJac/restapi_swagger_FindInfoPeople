package api

import (
	"encoding/json"
	// "fmt"
	swapi "main/generate"
	"net/http"
	"sync"
)

type PetStore struct {
	Pets   map[int64]PetStore
	NextId int64
	Lock   sync.Mutex
}

var _ swapi.ServerInterface = (*PetStore)(nil)

func NewPetStore() *PetStore { // cjplftv <L
	return &PetStore{
		Pets:   make(map[int64]Pet),
		NextId: 1000,
	}
}

// Хендлер
// sendPetStoreError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendPetStoreError(w http.ResponseWriter, code int, message string) { // обработака ошибок
	petErr := Error{
		Code:    int32(code),
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(petErr)
}

func (p *PetStore) FindPets(w http.ResponseWriter, r *http.Request, params FindPetsParams) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	var result []Pet

	for _, pet := range p.Pets {
		if params.Tags != nil {
			// If we have tags,  filter pets by tag
			for _, t := range *params.Tags {
				if pet.Tag != nil && (*pet.Tag == t) {
					result = append(result, pet)
				}
			}
		} else {
			// Add all pets if we're not filtering
			result = append(result, pet)
		}

		if params.Limit != nil {
			l := int(*params.Limit)
			if len(result) >= l {
				// We're at the limit
				break
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}
