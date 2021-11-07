import React from 'react';
import Item from './Item.js'


class Home extends React.Component {
    
    constructor(props) {
        super(props);
        this.state = {
            imgArray: []
        };
        this.selectItem = this.selectItem.bind(this);

    }

    componentDidMount() {
        // TODO: Make API call to get products
        this.setState({
            imgArray: [{name: "Pizza", price: 5.99, link:'pizza_img.png'},
            {name:"Ice cream", price:2.99, link: 'ice_cream_img.png'}]
        });
    }


    render(){
        let imageArrayRender = this.state.imgArray.map(image => {
            return (
                <Item 
                    name={image.name}
                    price={image.price}
                    selectItem={this.selectItem}
                    link={image.link}
                />
            );
        });
        return (
            <>
                {imageArrayRender}
            </>
        );
    }

    selectItem = (itemName) => {
        console.log(itemName);
    }
}


export default Home;