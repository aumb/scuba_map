package initializers

import (
	"os"

	supa "github.com/nedpals/supabase-go"
)

var DB *supa.Client

func ConnectToDB() {
	supabaseUrl := os.Getenv("DATABASE_URL")
	supabaseKey := os.Getenv("DATABASE_KEY")
	DB = supa.CreateClient(supabaseUrl, supabaseKey)
}
