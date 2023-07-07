package booking

import (
	"BookingService/src/model"
	"BookingService/src/module/seatBooked"
	"BookingService/src/module/user"
	SEATS "BookingService/src/module/seats"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(emailOrPhone string) (model.Bookings, error)
	FindAllBookings(user model.User) ([]model.Booking, error)
	FindAllBookingsSeats(bookings []model.Booking) (model.Bookings, error)
	Create(booking model.CreateBooking) (model.Booking, error)
	CountRows() (int64)
	AddToDB(booking model.Booking) (model.Booking, error)
}

type repository struct {
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) AddToDB(bookSeats model.Booking) (model.Booking, error){
	if res := r.client.Create(&bookSeats); res.Error != nil {
		return model.Booking{}, res.Error
	}

	return bookSeats, nil
}

func (r *repository) CountRows() (int64){
	var count int64
	_ = r.client.Model(&model.Booking{}).Count(&count)
	return count
}

func (r *repository) FindAllBookingsSeats(bookingsByUser []model.Booking) (model.Bookings, error){
	var bookings model.Bookings
	for _, bookingByUser := range bookingsByUser {
		seats, err := seatBooked.InitRepository(r.client).FindAllByBookingId(bookingByUser)
		if err != nil{
			return model.Bookings{}, err
		}
		seatsBooked := []model.SeatPrice{}
		for _, seat := range seats {
			seatClass, err := SEATS.InitRepository(r.client).FindSeatClass(seat.SeatId)
			if err != nil{
				return model.Bookings{}, err
			}
			seatBooked := createSeatPrice(seat,seatClass)
			seatsBooked = append(seatsBooked, seatBooked)
		}

		booking := createBookingResp(bookingByUser, seatsBooked)
		bookings.Bookings = append(bookings.Bookings, booking)
	}

	return bookings, nil
}

func (r *repository) FindAllBookings(user model.User) ([]model.Booking, error) {
	var bookingsByUser []model.Booking
	if res := r.client.Where(
		"user_id=?", user.ID).Find(&bookingsByUser); res.Error != nil {
		return []model.Booking{}, res.Error
	}
	return bookingsByUser,nil
}

func (r *repository) FindAll(emailOrPhone string) (model.Bookings, error) {
	email := emailOrPhone
	phone, _ := strconv.Atoi(emailOrPhone)
	//find uid
	user, err := user.InitRepository(r.client).FindViaEmailOrPhone(email, phone)
	if err != nil{
		return model.Bookings{}, err
	}

	//find all bookings
	bookingsByUser,err := r.FindAllBookings(user)
	if err != nil{
		return model.Bookings{}, err
	}

	//find all seats of the booking
	bookings, err := r.FindAllBookingsSeats(bookingsByUser)
	if err != nil{
		return model.Bookings{}, err
	}

	//return the bookings
	return bookings, nil
}

func (r *repository) Create(booking model.CreateBooking) (model.Booking, error) {
	isAnyBooked, err := seatBooked.InitRepository(r.client).CheckBookingSeat(booking.Seats)
	if isAnyBooked {
		return model.Booking{}, errors.New("seat already booked" + err.Error())
	}

	//valid user
	user, err := user.InitRepository(r.client).Find(user.User{
		Name: booking.Name,
		Phone: booking.Phone,
	})
	if err != nil{
		return model.Booking{}, err
	}

	//cal total
	total := calSeatBookingTotal(booking.Seats)

	//booking rows
	var count int64 = r.CountRows()
	bookSeats := createBooking(count,user,booking,total)

	//add to db
	bookSeats, err = r.AddToDB(bookSeats)
	if err!= nil{
		return model.Booking{}, err
	}

	_, err = seatBooked.InitRepository(r.client).Create(bookSeats, booking.Seats)
	if err!= nil{
		return model.Booking{}, err
	}

	//return booking
	return bookSeats, nil
}
