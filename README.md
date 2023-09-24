
# SKILLS TEST ABET JEAN-BAPTISTE#

  

### Requirements ###
* [Docker]

### Launch the project ###
 + Complete de docker-compose.yml environment variable
 + Build and launch the project with docker-compose
	+ `docker-compose build && docker-compose up -d`
	
### Use the project ###
 You can make request to this api on the route /analysis with parameter "duration" and "dimension":
 + duration accept time at format second (s), minute (m), and hour (h). Example of valide param: 5s, 10m, 7h
 + dimension accept value: 
	 + like
	 + comments
	 + favorites
	 + retweets
	 
The api return a json with this format :
 ```json
{
	"total_post": ,
	"minimum_timestamp": ,
	"maximum_timestamp": ,
	"p50": ,
	"p90": ,
	"p99": ,
}
```
 

### Architecture ###
The api is constituted of 4 layers of files:

 1. main (where the server is init)
 2. routes (where the routes are declared)
 3. controllers ( to verify the parameters) 
 4.  manager to execute de job logic.

  
The objects needed for the project are in the folder "models". There is one interface for the sseclient to switch between the real client and a stub needed for the tests

No libraries are used for this project so the go.mod is empty.

### Efficiency ###

 During the analysis execution, only the analysis struct and a slice of int are stored instead of all the events making this project quite lightweight. 
For more efficiency we could imagine an architecture where only one sseclient catch the event from the sse server then broadcast to the analysis needing them

### Scalability ###
Scale lineary with the number of client.
Could be optimized with one sseclient to put less presure on the sseserver


### Other feature ###
For a enterprise project I'll have added a swagger 