import axios, { formToJSON } from 'axios';

const API_BASE_URL = 'http://localhost:8080/api/ride';

//TODO noch nicht ausgereift
// nur ne Idee

const RideService = {

  async addRide(id, from, to) {

    console.log("startLat: ", from.lat, "startLon: ", from.lng);
    console.log("endLat: ", to.lat, "endLon: ", to.lng);
    try {
      const response = await axios.post(`${API_BASE_URL}/request`, {
        customer_id: id,
        start_location: {
          lat: from.lat,
          lon: from.lng,
        },
        end_location: {
          lat: to.lat,
          lon: to.lng
        }
      });

    console.log(response);

      console.log(response.data);

      return 'booking successful';

    } catch (error) {

      console.log(error.response.data);

      if (error.response.status === 404) {
        return 'no driver found';
      }
      else if (error.response.status === 400) {
        return 'bad request';
      }
      else {
        return 'internal server error oooops';
      }
    }
  },

  async getRideByDriverID(id) {
    try {
      const response = await axios.get(`${API_BASE_URL}/driver/${id}`);
      return response.data.rides;
    } catch (error) {
      console.error(error);
    }
  },
  
  async getRideByCustomerID(id) {
    try {
      const response =  await axios.get(`${API_BASE_URL}/customer/${id}`);
      return response.data.rides
    } catch (error) {
      console.error(error);
    }
  }
}

export default RideService;