import React from "react";
import SearchBar from "../searchbar/searchbar";

import "./styles.scss";
// https://www.inkbucket.com.ar/
const Header = () => {
  return (
    <header className="header">
      <div className="header-container">
        <div className="header-logo">
          <h3>GeekStore</h3>
        </div>
        <div className="header-navbar">
          <ul className='="header-navbar__list'>
            <li>Home</li>
            <li>Categorías</li>
            <li>Productos</li>
            <li>Personalizado</li>
          </ul>
        </div>
        <div className="header-searchbar">
          <SearchBar />
        </div>
      </div>
    </header>
  );
};

export default Header;
