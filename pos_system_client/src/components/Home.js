import React from 'react';
import Item from './Item.js'
class Home extends React.Component {
    
    constructor(props) {
        super(props);

        this.selectItem = this.selectItem.bind(this);

    }


    render() {
        return (
            <>
                <Item selectItem={this.selectItem}


                />
            </>
        );
    }

    selectItem = (itemName) => {
        console.log(itemName);
    }
}


export default Home;