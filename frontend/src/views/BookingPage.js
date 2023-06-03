import React, { useState, useEffect, useRef } from 'react';
import '../styles/BookingPage.css';
import MapComponent from '../components/MapComponent';
import RideService from '../services/RideService';
import User from '../model/user';
import Address from '../model/address';
import MapboxAutocomplete from 'react-mapbox-autocomplete';

var accessToken = process.env.mapbox_access_token;
accessToken =
  'pk.eyJ1IjoidW5qYWNvIiwiYSI6ImNsZ21mMjZ1dzA1NGEzcXFteGNiaWE5MXUifQ.UiRiOn2xmCGBYMu0vkvcyw';

function Booking({ activeUser }) {
  
  const [addressFrom, setAddressFrom] = useState(new Address());
  const [addressTo, setAddressTo] = useState(new Address());

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (addressFrom.lat !== undefined && addressTo.lat !== undefined) {
     const message = await RideService.addRide(activeUser.id, addressFrom, addressTo);
      alert(message);
    } else {
      alert('Please fill in all fields!');
    }
  };

  function _suggestionSelectFrom(result, lat, long, text) {
    setAddressFrom(new Address(parseFloat(lat), parseFloat(long)));
    console.log(result, "lat: ", lat, "lon: ", long, text);
    console.log(addressFrom);
  }

  function _suggestionSelectTo(result, lat, long, text) {
    setAddressTo(new Address(parseFloat(lat), parseFloat(long)));
    console.log(result, lat, long, text);
    console.log(addressTo);
  }

  return (
    <div className="booking-container">
      <div className='booking-page'>
        <div className="booking-content">
          <form className="inputForm">
            <MapboxAutocomplete
              publicKey={accessToken}
              inputClass="form-control search"
              onSuggestionSelect={_suggestionSelectFrom}
              country="de"
              resetSearch={false}
              placeholder="From"
            />
          </form>
          <form className="inputForm">
            <MapboxAutocomplete
              publicKey={accessToken}
              inputClass="form-control search"
              onSuggestionSelect={_suggestionSelectTo}
              country="de"
              resetSearch={false}
              placeholder="To"
            />
          </form>

          <div className="button-container">
            <button onClick={handleSubmit}>Book</button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Booking;
