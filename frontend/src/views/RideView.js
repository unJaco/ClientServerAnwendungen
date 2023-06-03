import '../styles/RideView.css';
import React from 'react';
import RideComponent from '../components/RideComponent';

function RideView({takenRides, drivenRides}) {

  console.log(takenRides);
  console.log(drivenRides);

  return (
    <div className="ride-view-page">
      <RideComponent takenRides={takenRides} drivenRides={drivenRides} />
    </div>

  );
}

export default RideView;
