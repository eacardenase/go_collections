package maps

import "fmt"

func main() {
	websites := map[string]string{
		"gcp": "https://cloud.google.com",
		"aws": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["aws"])

	websites["azure"] = "https://portal.azure.com"
	fmt.Println(websites)

	delete(websites, "gcp")
	fmt.Println(websites)
}
