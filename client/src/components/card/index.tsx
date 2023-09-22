import React from 'react';
import { Link } from "react-router-dom";
import type { CardProps } from '@types';

const Card = ({
    title,
    description,
    img,
    link,
    onClose,
}: CardProps ) => {
    return (
        <div className="card">
            <div className="card-main">
                <Link to={link}>
                    <div className="card-main__img">
                        <img src={img} alt="card-img"/>
                    </div>
                    <div className="card-main__content">
                        <h3>{title}</h3>
                        <p>{description}</p>
                    </div>
                </Link>
            </div>
            {
                onClose 
                && <div className="card-close">
                    <button className="card-close__button">
                        X
                    </button>
                </div>
            }
        </div>
    )
}

export default Card;