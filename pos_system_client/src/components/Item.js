import React from 'react';
import picture from '../pictures/logo512.png'
import './Item.css';

class Item extends React.Component {

    // constructor(props) {
    //     super(props);

    // }
    render() {
        const {name, price, selectItem, link} = this.props;

        return (
            <div className="product">
                <img className="product-picture" src={link} alt="Logo" onClick={() => selectItem("React Logo click!!")}/>
                <h6>{name}</h6>
                <h6>{price}</h6>
            </div>
        );
    }
}


export default Item;