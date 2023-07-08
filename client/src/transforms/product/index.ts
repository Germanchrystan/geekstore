type Product = {
    id: number;
    name: string;
    price: number;
    img: string;
    discount: {
        percentage: number;
        discounterPrice: number;
        finishDate: number;
    }
}

export const TransformToProductList = (products: Product[]) => {
    return products.map((p) => ({
        key: p.id,
        title: p.name,
        description: p.discount.percentage > 0  
        ? String(p.discount.discounterPrice) 
        : String(p.price),
        img: p.img,
    }))
}