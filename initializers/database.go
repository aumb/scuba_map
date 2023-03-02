package initializers

import (
	"os"

	"github.com/aumb/scuba_map/supabase"
)

var Client *supabase.Client

func ConnectToDB() {
	supabaseUrl := os.Getenv("DATABASE_URL")
	supabaseKey := os.Getenv("DATABASE_KEY")
	Client = supabase.CreateClient(supabaseUrl, supabaseKey)
}
