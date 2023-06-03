import React from 'react';
import '../styles/RideComponent.css';

const RideComponent = ({ takenRides, drivenRides }) => (
    <div className="ride-view">
        {(takenRides !== undefined && takenRides.length > 0) || (drivenRides !== undefined && drivenRides.length > 0) ? (
            <>
                {takenRides !== undefined && takenRides.length > 0 && (
                    <div className="ride-container">
                        <h2>Rides you have taken</h2>
                        <div className="ride-table">
                            <table>
                                <thead>
                                    <tr>
                                        <th>Driver</th>
                                        <th>Customer</th>
                                        <th>Distance</th>
                                        <th>Status</th>
                                        <th>Price</th>
                                        <th>Rating</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {takenRides.map((ride, index) => (
                                        <tr key={index}>
                                            <td>{ride.driverId}</td>
                                            <td>{ride.customerId}</td>
                                            <td>{ride.distance.toFixed(0)} km</td>
                                            <td>{ride.status}</td>
                                            <td>{ride.price.toFixed(2)}</td>
                                            <td>{ride.rating}</td>
                                        </tr>
                                    ))}
                                </tbody>
                            </table>
                        </div>
                    </div>
                )}

                {drivenRides !== undefined && drivenRides.length > 0 && (
                    <div className="ride-container">
                        <h2>Rides you have driven</h2>
                        <div className="ride-table">
                            <table>
                                <thead>
                                    <tr>
                                        <th>Driver</th>
                                        <th>Customer</th>
                                        <th>Distance</th>
                                        <th>Status</th>
                                        <th>Price</th>
                                        <th>Rating</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {drivenRides.map((ride, index) => (
                                        <tr key={index}>
                                            <td>{ride.driverId}</td>
                                            <td>{ride.customerId}</td>
                                            <td>{ride.distance.toFixed(0)} km</td>
                                            <td>{ride.status}</td>
                                            <td>{ride.price.toFixed(2)}</td>
                                            <td>{ride.rating}</td>
                                        </tr>
                                    ))}
                                </tbody>
                            </table>
                        </div>
                    </div>
                )}
            </>
        ) : (
            <div>
                <h2>No rides found</h2>
                <p>You have no rides in the register.</p>
            </div>
        )}
    </div>
);

export default RideComponent;


