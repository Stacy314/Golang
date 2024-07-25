package passwordmanager

import (
	"encoding/json"
	"errors"
	"os"
)

type Manager struct {
	filePath  string
	passwords map[string]string
}

func NewManager(filePath string) (*Manager, error) {
	manager := &Manager{
		filePath:  filePath,
		passwords: make(map[string]string),
	}

	if err := manager.load(); err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) load() error {
	file, err := os.Open(m.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Немає файлу, значить починаємо з порожнього словника
		}
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&m.passwords)
}

func (m *Manager) save() error {
	file, err := os.Create(m.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(m.passwords)
}

func (m *Manager) ListNames() []string {
	names := make([]string, 0, len(m.passwords))
	for name := range m.passwords {
		names = append(names, name)
	}
	return names
}

func (m *Manager) SavePassword(name, password string) error {
	m.passwords[name] = password
	return m.save()
}

func (m *Manager) GetPassword(name string) (string, error) {
	password, exists := m.passwords[name]
	if !exists {
		return "", errors.New("пароль не знайдено")
	}
	return password, nil
}
