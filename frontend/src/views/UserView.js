// src/components/LoggedInUser.js

import React, { useState, useEffect } from 'react';
import User from '../model/user';
import '../styles/UserView.css';

const UserView = ({activeUser}) => {
  const [user, setUser] = useState(new User());

  useEffect(() => {
   

    if (activeUser) {
      console.log(activeUser)
      setUser(activeUser);
    }
  }, [activeUser]);



  if (!user) {
    return <div className='user-view'>Loading user data...</div>;
  }

  return (
    <div className='user-view'>
      <h2>Logged In User</h2>
      <p>First Name: {user.firstName}</p>
      <p>Last Name: {user.lastName}</p>
      <p>Email: {user.email}</p>
      <p>Phone: {user.telNr}</p>
      <img src={user.profileUrl} alt={`${'No Picture'}`} />
    </div>
  );
};

export default UserView;
