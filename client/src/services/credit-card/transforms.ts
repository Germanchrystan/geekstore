import type { CreditCardList, CardProps } from '@types'
 
export const TransformToCardList = (cardList: CreditCardList[]): CardProps[] => {
    return cardList.map((c) => ({
        title: `...${c.last_code_number}`,
        description: c.company,
        img: '',
        link: `user/cards/${c.id}`

    }))
}