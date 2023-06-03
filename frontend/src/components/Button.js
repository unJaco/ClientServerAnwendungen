import React from 'react';
import '../styles/Button.css';

function Button(props) {
  return (
    <button className="btn" onClick={props.onClick}>
      {props.text}
    </button>
  );
}

export default Button;
