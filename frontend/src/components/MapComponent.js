import React, { useEffect, useRef, useState } from 'react';
import mapboxgl from 'mapbox-gl';
import '../styles/MapComponent.css';
import Address from '../model/address';

// Replace with your own MapBox access token
var accessToken = process.env.mapbox_access_token;
accessToken = "pk.eyJ1IjoidW5qYWNvIiwiYSI6ImNsZ21mMjZ1dzA1NGEzcXFteGNiaWE5MXUifQ.UiRiOn2xmCGBYMu0vkvcyw";

const MapComponent = ({ from, to }) => {
  const mapContainer = useRef(null);
  const map = useRef(null);
  const [lng, setLng] = useState(-74.006);
  const [lat, setLat] = useState(40.7128);
  const [zoom, setZoom] = useState(9);

  useEffect(() => {
    if (map.current) {
      map.current.setCenter([lng, lat]);
      return;
    }

    map.current = new mapboxgl.Map({
      container: mapContainer.current,
      style: 'mapbox://styles/mapbox/streets-v11',
      center: [lng, lat],
      zoom: zoom,
      accessToken: accessToken
    });

    map.current.on('move', () => {
      setLng(map.current.getCenter().lng.toFixed(20));
      setLat(map.current.getCenter().lat.toFixed(20));
      setZoom(map.current.getZoom().toFixed(2));
    });

    if (from && to) {
      const fromCoordinates = new Address(from.lng, from.lat);
      const toCoordinates = new Address(to.lng, to.lat);
      console.log('1')
      map.current.on('load', async () => {
        map.current.addSource('route', {
          type: 'geojson',
          data: {
            type: 'Feature',
            properties: {},
            geometry: {
              type: 'LineString',
              coordinates: [fromCoordinates, toCoordinates]
            }
          }
        });
console.log('mitte if');
        map.current.addLayer({
          id: 'route',
          type: 'line',
          source: 'route',
          layout: {
            'line-join': 'round',
            'line-cap': 'round'
          },
          paint: {
            'line-color': '#888',
            'line-width': 8
          }
        });
        // Fit the map to the route bounds
        const bounds = new mapboxgl.LngLatBounds();
        bounds.extend(fromCoordinates);
        bounds.extend(toCoordinates);
        map.current.fitBounds(bounds, { padding: 50 });

        // Jump to the starting point of the route
        map.current.jumpTo({ center: fromCoordinates });
      });
    }
  }, [lng, lat, zoom, from, to]);

  return (
    <div className="map-component">
      <div className="map-container" ref={mapContainer} />
    </div>
  );
};

export default MapComponent;
