### k8stestbed

This repo will serve as the home of personal messing around with kubernetes, especially advanced concepts within cluster (and other) operations

The general plan for this project is to set up a minimal cluster within GKE, and create various microsercices linked via service mesh / grpc. 

This will also serve as a testbed for unguided / non-toutorial based learing of golang. 


## Project Idea:

In need of a basic idea for a set of microservices, I'm settling on:
* Scrape baseball-reference.com (or something similar) for stats on a single player (for now)
* Store this data (initially in memory, but when expanding, will need a database?)
* display this data in a web-app (RBIs per year graph or something)


## Project plan:

* Stand up GKE cluster (dont need anything too crazy here)
* Helm charts for templatization
* CI/CD pipeline:
    * CircleCI
    * All work will go through this
    * its going to end up being pretty complicated, as I'm working on a mono-repo concept here.
    * Probably NOT use the box-method where you have a seperate repo of 'current state', although that could be added later
* maybe fool around with different contianer runtimes?
* Minikube for local dev?
* Microservices:
    * write in GO, 
    * package into from scratch docker contianers
    * Start with a Hello World webservice, proof of concepting
    * Branch out from there
    * add microservices that the helloworld endpoint will reference to get more information
        * time checking
        * picking a color - have the front-end grab this a bunch of times for a single 'call', so that we can create an array/matrix of the results
* Use these as a basis of making the cluster/pipeline/workflow better
    * Service Mesh
        * in-transit encryption
        * cannary deployments
    * Monitoring
    * properly size the pods


* CI pipleine: include squeeze/load testing, put results into PR so that 'devs' can see impact & pass / fail without loggin into another app
*  maybe try out some machine learning somewhere (way down the line...)
        

* maybe try some machine learning later?
* put results of CI pipeline / load testing into the PR lines, so that 'developers' can see differences without loggin into another tool