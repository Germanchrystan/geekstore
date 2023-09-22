import React, { useEffect, useState } from 'react';
import Button from '../../components/button';
import getPanelData from './helpers/get-panel-data';
import panelOptions from './constants';
import './styles.scss';


const Profile = () => {
    const [selectedPanel, setSelectedPanel] = useState('purchases');
    
    useEffect(() => {
        getPanelData(selectedPanel, 1);
    }, [selectedPanel]);

    return (
        <div className='profile'>
            <div className='profile-upper-card'>
                <div className='profile-upper-card__profile-pic'>
                    <img alt="profile-pic" src="https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_640.png" />
                </div>
                <div>
                    <h3>
                        Username
                    </h3>
                </div>
            </div>
            <div className='panel'>

                <div className='panel-list'>
                    {
                        panelOptions.map((option) =>
                            <Button
                                key={option}
                                onClick={() => setSelectedPanel(option)}
                                text={option}
                                disabled={selectedPanel === option}
                                size='small'
                            />)
                    }
                </div>
                <div className='panel-content'>

                </div>
            </div>
        </div>
    )
}

export default Profile;