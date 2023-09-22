export const TransformToPurchaseList = (purchaseList: PurchaseCard[]): CardProps[] => {
    return purchaseList.map(p =>( {
        title: String(p.total),
        description: p.address_name,
        img: '', //getPurchaseStateImage(p.status),
        link: `/purchase/${p.id}`
    }));
}