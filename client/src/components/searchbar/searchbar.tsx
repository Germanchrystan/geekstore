import React, { useState } from 'react';

const SearchBar = () => {
    const [input, setInput] = useState('');

    const onChangeInput = (e: any) => { // TODO change type
        const { target } = e;
        const { value } = target;
        setInput(value);
    }

    return(
        <input type='text' value={input} onChange={onChangeInput} />
    )
}

export default SearchBar;