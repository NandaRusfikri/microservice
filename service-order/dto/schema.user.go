package dto

type SchemaUser struct {
	Name      string ` json:"name,omitempty" validate:"required,lowercase"`
	Phone     string ` json:"phone,omitempty"`
	Address   string ` json:"address,omitempty"`
	PhotoPath string ` json:"photo_path,omitempty"`
	Status    string ` json:"status"`
	Username  string ` json:"username,omitempty"`
	Email     string ` json:"email,omitempty" validate:"required,email"`
	Password  string ` json:"password,omitempty"`
	Token     string ` json:"token,omitempty"`
}
