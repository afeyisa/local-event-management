export const filterDay = (day) => {
  const today = new Date()
  let startDate, endDate
  const currentDay = today.getDay()
  switch (day) {
  /* eslint-disable */
      case 'today':
        startDate = new Date(today)
        endDate = new Date(today)
        break
      case 'tomorrow':
        startDate = new Date(today)
        startDate.setDate(today.getDate() + 1)
        endDate = new Date(startDate)
        break

      case 'thisWeekend':
        startDate = new Date(today)
        startDate.setDate(today.getDate() + (5 - currentDay)) // Friday
        endDate = new Date(startDate)
        endDate.setDate(startDate.getDate() + 2) // Sunday
        break

      case 'nextWeek':
        startDate = new Date(today)
        startDate.setDate(today.getDate() + (8 - currentDay)) // Next Monday
        endDate = new Date(startDate)
        endDate.setDate(startDate.getDate() + 6) // Following Sunday
        break

      case 'nextMonth':
        startDate = new Date(today)
        endDate = new Date(today)
        endDate.setDate(today.getDate() + 30)
        break

      default:
        startDate = new Date(today)
        endDate = new Date(today)
      }

      return {
        _and: [
          { event_date: { _gte: startDate } },
          { event_date: { _lte: endDate } },
        ],
      }
    }
  