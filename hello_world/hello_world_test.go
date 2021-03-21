package hello_world

import "testing"

func TestHello(t *testing.T){

	assertCorrectMessage := func(t testing.TB, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("should return 'Hello World' string", func (t *testing.T){
		got := HelloWorld()
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
	
	t.Run("should return 'Hello Italo' string", func (t *testing.T){
		got := Hello("Italo")
		want := "Hello Italo"
	
		assertCorrectMessage(t, got, want)
	})
}