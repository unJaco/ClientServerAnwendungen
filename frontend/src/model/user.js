class User {
    constructor(id, firstName, lastName, profileUrl, telNr, email, pwHash) {
      this.id = id;
      this.firstName = firstName;
      this.lastName = lastName;
      this.profileUrl = profileUrl;
      this.telNr = telNr;
      this.email = email; 
      this.pwHash = pwHash;
    }
  }
  
  export default User;
  