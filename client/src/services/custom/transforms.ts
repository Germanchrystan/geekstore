import type { CustomProductCard, CardProps } from '@types'

export const TransformToCustomList = (customList: CustomProductCard[]): CardProps[] => {
    return customList.map((c) => ({
        title: c.name,
        // TODO: display status
        description: String(c.price),
        img: '',
        link: `/product/custom/${c.id}`
    }))
}