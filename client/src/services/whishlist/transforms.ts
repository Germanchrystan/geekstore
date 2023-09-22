export const TransformToProductList = (products: ProductCard[]): CardProps[] => {
    return products.map((p) => ({
        key: p.id,
        title: p.name,
        description: p.discount.percentage > 0  
        ? String(p.discount.discounterPrice) 
        : String(p.price),
        img: p.img,
        link: `/product/${p.id}`
    }))
}