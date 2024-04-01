package main

import "fmt"

type Offer struct {
	RewarId    int
	RewardType string
}

type Restaurant struct {
	Name   string
	Offers []Offer
}

func main() {
	offer1 := Offer{RewarId: 123, RewardType: "weekly"}
	offer2 := Offer{RewarId: 456, RewardType: "monthly"}
	offer3 := Offer{RewarId: 341, RewardType: "rental"}
	offer4 := Offer{RewarId: 876, RewardType: "daily"}
	offer5 := Offer{RewarId: 900, RewardType: "rental"}

	r := Restaurant{
		Name:   "KFC",
		Offers: []Offer{offer1, offer2, offer3, offer4, offer5},
	}
	fmt.Println("Restaurant name:", r.Name)
	fmt.Println("Offers:")
	for _, offer := range r.Offers {
		fmt.Printf("Reward ID: %d, Reward type: %s\n", offer.RewarId, offer.RewardType)
	}

	newSlice := make([]int, 0)
	for _, offer := range r.Offers {
		if offer.RewardType == "rental" {
			newSlice = append(newSlice, offer.RewarId)
		}
	}
	fmt.Println("Reward Id of offers with type 'rental':", newSlice)

	offer6 := Offer{RewarId: 12, RewardType: "weekly"}
	offer7 := Offer{RewarId: 17, RewardType: "rental"}
	offer8 := Offer{RewarId: 19, RewardType: "monthly"}
	r2 := Restaurant{
		Name:   "Subway",
		Offers: []Offer{offer6, offer7, offer8},
	}
	//made a list of restaurants having two objects
	restaurants := []Restaurant{r, r2}
	newSlice2 := getRentalRewardIds(restaurants)

	fmt.Println("Reward Ids of offers with type 'rental' from the list of restaurant is:", newSlice2)
}

//function to get reward ids of restaurants with offer of type rental
func getRentalRewardIds(restaurants []Restaurant) []int {
	newSlice2 := make([]int, 0)

	for _, restaurant := range restaurants {
		for _, offer := range restaurant.Offers {
			if offer.RewardType == "rental" {
				newSlice2 = append(newSlice2, offer.RewarId)
			}
		}
	}
	return newSlice2
}
