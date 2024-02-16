package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FirstName     string             `json:"first_name" bson:"first_name,omitempty"`
	LastName      string             `json:"last_name" bson:"last_name,omitempty"`
	UserName      string             `json:"username" bson:"username,omitempty"`
	Email         string             `json:"email" bson:"email,omitempty"`
	Password      string             `json:"password" bson:"password,omitempty"`
	Image         string             `json:"image" bson:"image"`
	InvitedCode   string             `json:"invite_code" bson:"invite_code"`
	ProviderCode  string             `json:"provider_code" bson:"provider_code"`
	MainBalance   string             `json:"main_balance" bson:"main_balance"`
	DemoBalance   string             `json:"demo_balance" bson:"dem_balance"`
	LiveBalance   string             `json:"live_balance" bson:"live_balance"`
	Role          string             `json:"role" bson:"role"`
	OtpCode       string             `json:"otp_code" bson:"otp_code"`
	HasTwoFa      bool               `json:"has_two_fa" bson:"has_two_fa"`
	EmailVerified bool               `json:"email_verified" bson:"email_verified"`
	IsVerify      bool               `json:"is_verify" bson:"is_verify"`
	LoginCount    int                `json:"login_count" bson:"login_count"`
	LastLogin     time.Time          `json:"last_login" bson:"last_login"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
