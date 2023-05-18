import React from 'react';
import SearchBar from '../searchbar/searchbar';

const Header = () => {
    return (
    <header className="header">
        <div className="header-logo">

        </div>
        <div className="header-searchbar">
            <SearchBar />
        </div>
        <div className="header-navbar">
            <ul className='="header-navbar__list'>
                <li></li>
            </ul>
        </div>
    </header>
    )
}

export default Header;