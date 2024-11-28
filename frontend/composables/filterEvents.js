import { filterDay } from './filterDays'

export const constructWhereClause = (keywords) => {
  return {
    _or: [
      { _or: keywords.map(word => ({ title: { _ilike: word } })) },
      { _or: keywords.map(word => ({ description: { _ilike: word } })) },
      { _or: keywords.map(word => ({ _and: { event_tags: { _and: { tag: { _and: { tag_word: { _ilike: word } } } } } } })) },
      { _or: keywords.map(word => ({ category: { category_name: { _ilike: word } } })) },
    ] }
}

export const filteredWhereClause = (selecteDays, selectedCategories, searchTerm, isFree) => {
  let filtereddDays = [{ event_date: { } }]

  if (selecteDays.length > 0) {
    filtereddDays = selecteDays.map(day => filterDay(day))
  }

  let filteredCatagoryIds = [{ category_id: { } }]
  if (selectedCategories.length > 0) {
    filteredCatagoryIds = selectedCategories.map(id => ({ category_id: { _eq: id } }))
  }
  let ticketPrice = {}
  if (isFree === 'true') {
    ticketPrice = { _eq: 0 }
  }
  if (isFree === 'false') {
    ticketPrice = { _gt: 0 }
  }
  const keywords = searchTerm.split(' ').map(word => `%${word}%`)

  const whereClause = {
    _and: [
      constructWhereClause(keywords),
      {
        ticket_price: ticketPrice,
      },
      {
        _or: filteredCatagoryIds,
      },
      {
        _or: filtereddDays,
      },
    ],
  }
  return whereClause
}

export const filteredWhereClauseWithLocation = (selecteDays, selectedCategories, isFree) => {
  let filtereddDays = [{ event_date: { } }]

  if (selecteDays.length > 0) {
    filtereddDays = selecteDays.map(day => filterDay(day))
  }

  let filteredCatagoryIds = [{ category_id: { } }]
  if (selectedCategories.length > 0) {
    filteredCatagoryIds = selectedCategories.map(id => ({ category_id: { _eq: id } }))
  }
  let ticketPrice = {}
  if (isFree === 'true') {
    ticketPrice = { _eq: 0 }
  }
  if (isFree === 'false') {
    ticketPrice = { _gt: 0 }
  }

  const whereClause = {
    _and: [
      {
        ticket_price: ticketPrice,
      },
      {
        _or: filteredCatagoryIds,
      },
      {
        _or: filtereddDays,
      },
    ],
  }
  return whereClause
}
