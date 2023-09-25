//package main

//import (
//	"fmt"
//)

/*func main() {
	var arr = []int{9, 8, 7, 4, 5, 3}
	fmt.Println("Unsorted array of strings is", arr)
	sort.Ints(arr)
	fmt.Println("The above array is sorted and the result is:", arr)
}*/

/*func sortArray(arr [5]int, min int, temp int) [5]int {
	for i := 0; i <= 4; i++ {
		min = i
		for j := i + 1; j <= 4; j++ {
			if arr[j] < arr[min] {

				min = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
	return arr
}
func main() {
	arr := [5]int{6, 4, 8, 2, 3}
	fmt.Println("The unsorted array is:", arr)
	var min int = 0
	var temp int = 0
	array := sortArray(arr, min, temp)
	fmt.Println()
	fmt.Println(" after sorting array is:", array)
}*/

package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WeatherData struct {
	//gorm.Model
	//WisId              uint      `json:"wis_id" gorm:"primary_key"`
	DeviceNum          int     `json:"device_num"  binding:"required"`
	DeviceName         string  `json:"device_name"  binding:"required"`
	DateTime           string  `json:"date_time" binding:"required"`
	WindSpeed          int     `json:"wind_speed"  binding:"required"`
	WindSpeed_Kmh      float64 `json:"wind_speed_kmh"  binding:"required"`
	WindDirection      float64 `json:"wind_direction"  binding:"required"`
	Windcardinal       float64 `json:"wind_cardinal"  binding:"required"`
	AirTemperature     float64 `json:"air_temperature"  binding:"required"`
	WeatherStation     string  `json:"weatherstation"  binding:"required"`
	Relativehumidity   float64 `json:"relative_humidity"  binding:"required"`
	AtmospherePressure float64 `json:"atmosphere_pressure"  binding:"required"`
	Visibility         float64 `json:"visibility"  binding:"required"`
	IrstPav            float64 `json:"irst_pav"  binding:"required"`
	BatteryValues      float64 `json:"battery_values"  binding:"required"`
	Rain               float64 `json:"rain"  binding:"required"`
	RoadTemperature    float64 `json:"road_temperature"  binding:"required"`
	Dateadded          string  `json:"date_added"  binding:"required"`
}

var DB *gorm.DB

func GetAllData(c *gin.Context) {
	var weather_data []WeatherData
	err := DB.Find(&weather_data).Error
	if err != nil {
		fmt.Println("err while fetch : ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": weather_data})
}

func ConnectDatabase() {
	//var hourly Hourly
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/itms"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&WeatherData{})

	DB = db

}

func main() {
	ConnectDatabase()
	r := gin.Default()

	r.GET("/weather_data", func(c *gin.Context) {
		//wdata :=WeatherData{}
		var weather_data []WeatherData

		err := DB.Find(&weather_data).Error
		if err != nil {
			fmt.Println("err while fetch : ", err)
			return
		}
		device_num := c.DefaultQuery("device_num", "4242")
		var weather_data_result []WeatherData
		for _, data := range weather_data {
			if strconv.Itoa(data.DeviceNum) == device_num {
				weather_data_result = append(weather_data_result, data)
			}
		}
		if device_num == "" {
			weather_data_result = weather_data
		}

		sort.Slice(weather_data_result, func(i, j int) bool {
			return weather_data_result[i].DateTime < weather_data_result[j].DateTime
		})

		//

		c.JSON(200, weather_data_result)
	})

	r.Run("localhost:8081")
}

//fmt.Println("Sorting an array of time")
//const shortForm = "2006-01-02 15:04:05"
//t1, _ := time.Parse(shortForm, "2023-08-17 14:31:00")
//t2, _ := time.Parse(shortForm, "2023-08-17 14:16:00")
//t3, _ := time.Parse(shortForm, "2023-08-17 14:17:00")
//t4, _ := time.Parse(shortForm, "2023-08-17 14:20:00")

//weather_data = []WeatherData

/*sql:="select * from weather_data"

if sort := c.Query(key:"sort"); sort != ""{
	sql = fmt.Sprintf(format: "%s ORDER BY date_time %s",sql,sort )
}

db.raw(sql).Scan(&weather_data)
return c.json(weather_data)
	})
	r.Run("localhost:8081")
}*/

/*		for _, wdata := range weather_data {
			fmt.Println(wdata)
		}
		sort.Sort(ByDateTime(weather_data))
		fmt.Println("sorted:")

		for _, wdata := range weather_data {
			fmt.Println(wdata)
		}
		c.JSON(http.StatusOK, gin.H{"message": "This is for weather_data", "data": weather_data})
	})
	//r.GET("/weather_data", GetAllData)
	r.Run("localhost:8081")
}

/*var weather_data_map = make(map[string]WeatherData)
	weather_data_map["1"] = WeatherData{date_time: time.Now().Add(12 * time.Hour)}
	weather_data_map["2"] = WeatherData{date_time: time.Now()}
	weather_data_map["3"] = WeatherData{date_time: time.Now().Add(24 * time.Hour)}
	//Sort the map by date
	date_time_data := make(ByDateTime, 0, len(weather_data_map))
	for _, d := range weather_data_map {
		date_time_data = append(date_time_data, d)
	}
	fmt.Println(date_time_data)
	sort.Sort(date_time_data)
	fmt.Println(date_time_data)
}*/

//c.JSON(http.StatusOK, gin.H{"message": "This is for weather_data", "data": weather_data})
//})
//router.Run(":8080")

//fmt.Println("Sorting an array of time")
//const shortForm = "2006-01-02 15:04:05"
//t1, _ := time.Parse(shortForm, "2023-08-17 14:31:00")
//t2, _ := time.Parse(shortForm, "2023-08-17 14:16:00")
//t3, _ := time.Parse(shortForm, "2023-08-17 14:17:00")
//t4, _ := time.Parse(shortForm, "2023-08-17 14:20:00")

//for _, t := range weather_data {
//fmt.Println(data)
