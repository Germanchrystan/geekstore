export const TransformToAddressList = (addresses: Address[]): CardProps[] => {
    return addresses.map((a) => ({
        title: a.name,
        description: `${a.street} ${a.street_number}, ${a.state} ${a.country}`,
        link: "",
        img: "",
    }));
}