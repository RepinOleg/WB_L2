package pattern

/*
Состояние — это поведенческий паттерн проектирования,
который позволяет объектам менять поведение в зависимости от своего состояния.
Извне создаётся впечатление, что изменился класс объекта.

Паттерн Состояние предлагает создать отдельные классы для каждого состояния,
в котором может пребывать объект, а затем вынести туда поведения, соответствующие этим состояниям.

Вместо того, чтобы хранить код всех состояний,
первоначальный объект, называемый контекстом, будет содержать ссылку на один из объектов-состояний
и делегировать ему работу, зависящую от состояния.

Очень важным нюансом, отличающим этот паттерн от Стратегии, является то, что и контекст,
и сами конкретные состояния могут знать друг о друге и инициировать переходы от одного состояния к другому.

*/

type MobileAlertStater interface {
	Alert() string
}

type MobileAlert struct {
	state MobileAlertStater
}

func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

type MobileAlertVibration struct {
}

func (a *MobileAlertVibration) Alert() string {
	return "Bzzz... Bzzz... Bzzzz..."
}

type MobileAlertSong struct {
}

func (a *MobileAlertSong) Alert() string {
	return "Some song"
}
