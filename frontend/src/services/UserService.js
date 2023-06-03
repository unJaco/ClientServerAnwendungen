import axios, { Axios, formToJSON } from 'axios';
import User from '../model/user';

const API_BASE_URL = 'http://localhost:8080/api/user';
const REG_BASE_URL = 'http://localhost:8080/api/user';
const Vehicles_BASE_URL = 'http://localhost:8080/api/vehicle/driverID'; 


const UserService = {

  async getAllUsers() {
    try {
      const response = await axios.get(`${API_BASE_URL}/all`);
      return response.data.map(user => new User(
        user.id,
        user.firstName,
        user.lastName,
        user.profileUrl,
        user.telNr,
        user.email
      ));
    } catch (error) {
      console.error('Error fetching users:', error);
      throw error;
    }
  },

  async getUserById(id) {
    try {
      const response = await axios.get(`${API_BASE_URL}/${id}`);
      const userData = response.data;
      return new User(
        userData.id,
        userData.firstName,
        userData.lastName,
        userData.profileUrl,
        userData.telNr,
        userData.email
      );
    } catch (error) {
      console.error(`Error fetching user with id ${id}:`, error);
      throw error;
    }
  },

  async getDriverById(driver_id) {
    try {
      const response = await axios.get(`${Vehicles_BASE_URL}/${driver_id}`);

      if(response != null){
        return this.getUserById(driver_id);
      }
      else{
        return null;
      }
    } catch (error) {
      console.error(`Error fetching user with id ${driver_id}:`, error);
      throw error;
    }
  },

  async login(email, pw_hash) {

    console.log(email)
    console.log(pw_hash)

    try {
      
      const response = await axios.post(`${API_BASE_URL}/login`, {
        params : {
          email : email, 
          pw_hash : pw_hash
        }  
       
      });

      const data = response.data;

      console.log(data);

      if (data.user != null) {
        return new User(
          data.user.ID,
          data.user.first_name,
          data.user.last_name,
          data.user.profile_url,
          data.user.tel_nr,
          data.user.email
        )
      } else {
        return data.message;
      }

    } catch (error) {
      alert('login failed')
      console.error(`Error with login`, error);
        throw error;
    }
  },

  async register(user){
    
    try {
      
      
      const response = await axios.post(`${REG_BASE_URL}/create`, {
       first_name : user.firstName,
       last_name : user.lastName,
       email : user.email,
       tel_nr : user.telNr,
       pw_hash : user.pwHash

      });

      const data = response.data;

      if (response.status === 500) {
          console.log(response)
      } else if (response.status === 400) {
        console.log(response)
      } else {
        return new User(
          data.user.ID,
          data.user.first_name,
          data.user.last_name,
          data.user.profile_url,
          data.user.tel_nr,
          data.user.email,
          data.user.pw_hash
        )
      }

    } catch (error) {
      console.error(`Error creating user`, error);
      throw error;
    }
  },
};

export default UserService;
