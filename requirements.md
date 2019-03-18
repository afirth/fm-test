### 1. Create GBDX server component
The component should accept a `POST` request containing GeoJSON, connect to GBDX and retrieve all matching records from the catalogue, returning GeoJSON to the client. The data returned should contain at a minimum properties representing the GBDX catalogID, platformName and timestamp.

The `aoi.geojson` file defines our area of interest and can be used to test the POST endpoint.

**Note:** You will need to create a free evaluation account on the GBDX platform.

### 2. Dockerise the component
Create `Dockerfile` and `docker-compose.yml` files for the component you have created in Step 1. Running `docker-compose up` should expose the service via a port on your local machine: you should be able to `POST` GeoJSON to the server via this port.

### 3. Deploy the component to a local Kubernetes cluster
Deploy the component to a Kubernetes instance that you have running on your local machine.

You should be able to:

* Access the service from outside of the cluster (i.e. be able to `POST` and receive GeoJSON from your local machine).
* Be able to scale up and down the number of instances of the component.

### 4. Document the deployment process
Please provide all the necessary files along with accompanying technical documentation on how to replicate the setup and deploy the cluster on a local machine.
