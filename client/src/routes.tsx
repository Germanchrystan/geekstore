import React, { useState } from 'react';
import { BrowserRouter, Route } from 'react-router-dom';
//==========================================================//
// Pages
//==========================================================//
import Home from './pages/home';

//==========================================================//
import Header from './components/header';
//==========================================================//

const Routes = () => {
    // const [user, setUser] = useState(() => {
    //     if(localStorage.getItem('ls_gd_userData') !== null){
    //         return JSON.parse(localStorage.getItem('ls_gd_userData')!)
    //     } else return null
    // })

    return (
        <BrowserRouter>
            <Route path="/*">
                <Header />
            </Route>
            <Route exact path="/">
                <Home />
            </Route>
            {/* <Route path="/game/:id">
                <GameDetail />
            </Route>
            <Route exact path="/login">
                <Login />
            </Route>
            <Route exact path="/register">
                <Register />
            </Route>
            <Route path="/">
                <Footer />
            </Route> */}
        </BrowserRouter>
    )
}

export default Routes;
