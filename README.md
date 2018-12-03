# GUC-SwitchHub
GUC SwitchHub automates the hassle of tutorial group switching between students at the GUC.

## Instructions to run the website

1. Make sure that Docker is installed and started on your machine.

2. Make sure that you have a working internet connection.

3. Navigate to the project folder using the terminal.

4. Use **docker-compose up -d --build** command to start the website and all the accompained services.

**Command Flags**:  
   **-d**: used to run the command in the daemon mode.  
   **--build**: used to ensure that the running version will be the latest build.

5. Use **cat database.sql | docker-compose exec -T db psql -U root GUCSwitchHubDB** command to reflect changes from the database.sql file to the GUCSwitchHubDB

**Command Flags**:  
**-U**: dbUser  
**GUCSwitcHubDB**: The dbName

6. Do another **docker-compose up -d --build** to ensure that everything is running as expected.


### Config Files

You are required to create a **JSON** config file, named **config.json**, place it in the project directory **on the same level as the main.go file**.

#### An example of the config.json file is provided below:

    {
	    "dbHost": "exampleHost",
	    "dbUser": "exampleUser",
	    "dbPassword": "examplePassword",
        "dbName": "exampleDbName"
    }
