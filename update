package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

func main() {
	router := gin.Default()

	router.GET("/update", updateUser)

	router.Run(":8080")
}

type User struct {
	Auth_id  string `json:"auth_id"`	
	Name  string `json:"name"`	
	Sex string `json:"sex"`
	Age string `json:"age"`
	Place string `json:"place"`
	Occupation string `json:occupation"`
	techs []string `json:"techs"`
	TopTechs []string `json:"topTechs"`
	Image_url string `json:" image_url"`
  Hobby string `json:"hobby"`
	Editor string `json:" editor"`
	Affiliation string `json:"affiliation"`
	Qualification []string `json:"qualification"`
	Message string `json:"message"`
	Portfolio string `json:"portfolio"`
	Graduate string `json:"graduate"`
	DesiredOccupation string `json:"desiredOccupation"`
	Faculty string `json:"faculty"`
	Experience []string `json:"experience"`
	Github string `json:"github"`
	Twitter string `json:"twitter"`
	Zenn string `json:"zenn"`
	Qiita string `json:"qiita"`
	Atcoder string `json:"atcoder"`
}

func updateUser(c *gin.Context) {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	client := supa.CreateClient(supabaseUrl, supabaseKey)

	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := req.ID
	if userID == "" {
		  c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
	}

	var inputData User
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser User
	err := client.DB.From("users").Select("name", "hobby").Eq("id", userID).Single(&existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedData := User{
		Name:  inputData.Name,  
		Hobby: inputData.Hobby, 
		Sex : inputData.Sex,
		Age : inputData.Age,
		Place : inputData.Place,
	  Occupation : inputData.Occupation,
		Techs : inputData.Techs,
		TopTechs : inputData.TopTechs,
		Image_url : inputData.Image_url,
		Editor : inputData.Editor,
		Affiliation: inputData.Affiliation,
		Qualification: inputData.Qualification,
		Message : inputData.Message,
		Portfolio : inputData.Portfolio,
		Graduate : inputData.Graduate,
		DesiredOccupation : inputData.DesiredOccupation,
		Faculty : inputData. Faculty,
		Experience: inputData.Experience,
		Github: inputData.Github,
		Twitter : inputData.Twitter,
		Zenn : inputData.Zenn,
		Qiita : inputData.Qiita,
		Atcoder : inputData.Atcoder,
	}

	if updatedData.Name == "" {
		updatedData.Name = existingUser.Name
	}
	if updatedData.Hobby == "" {
		updatedData.Hobby = existingUser.Hobby
	}
	if updatedData.Sex == "" {
		updatedData.Sex = existingUser.Sex
	}
	if updatedData.Age == "" {
		updatedData.Age = existingUser.Age
	}
	if updatedData.Place == "" {
		updatedData.Place = existingUser.Place
	}
	if updatedData.Occupation == "" {
		updatedData.Occupation = existingUser.Occupation
	}
	if updatedData.Techs == "" {
		updatedData.Techs = existingUser.Techs
	}
	if updatedData.TopTechs == "" {
		updatedData.TopTechs = existingUser.TopTechs
	}
	if updatedData.Image_url == "" {
		updatedData.Image_url = existingUser.Image_url
	}
	if updatedData.Editor == "" {
		updatedData.Editor = existingUser.Editor
	}
	if updatedData.Affiliation == "" {
		updatedData.Affiliation = existingUser.Affiliation
	}
	if updatedData.Qualification == "" {
		updatedData.Qualification = existingUser.Qualification
	}
	if updatedData.Message == "" {
		updatedData.Message = existingUser.Message
	}
	if updatedData.Portfolio == "" {
		updatedData.Portfolio = existingUser.Portfolio
	}
	if updatedData.Graduate == "" {
		updatedData.Graduate = existingUser.Graduate
	}
	if updatedData.DesiredOccupation == "" {
		updatedData.DesiredOccupation = existingUser.DesiredOccupation
	}
	if updatedData.Faculty == "" {
		updatedData.Faculty = existingUser.Faculty
	}
	if updatedData.Experience == "" {
		updatedData.Experience = existingUser.Experience
	}
	if updatedData.Github == "" {
		updatedData.Github = existingUser.Github
	}
	if updatedData.Twitter == "" {
		updatedData.Twitter = existingUser.Twitter
	}
	if updatedData.Zenn == "" {
		updatedData.Zenn = existingUser.Zenn
	}
	if updatedData.Qiita == "" {
		updatedData.Qiita = existingUser.Qiita
	}
	if updatedData.Atcoder == "" {
		updatedData.Atcoder = existingUser.Atcoder
	}
	
	var results map[string]interface{}
	err := client.DB.From("users").Update(updatedData).Eq("id", userID).Execute(&results)
	if err != nil {
		  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "results": results})
}
