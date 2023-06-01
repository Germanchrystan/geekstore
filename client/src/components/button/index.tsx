import React from 'react';
import './styles.scss';

type ButtonProps = {
    text: String,
    size: 'small' | 'medium' | 'big',
    color?: 'blue' | 'green' | 'pink' | 'yellow' 
    disabled: boolean | undefined,
    onClick: React.MouseEventHandler,
} 

const Button = ({
    text,
    color,
    disabled,
    size,
    onClick,
}: ButtonProps) => {

    return (
        <button
            className={`button button-${size} button-${color}`}
            disabled={disabled}
            onClick={onClick}
        >
            {text}
        </button>
    )
}

export default Button;