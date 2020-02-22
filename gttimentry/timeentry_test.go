package gttimeentry

import (
	"github.com/tumb1er/gtoggl-api/test"
	"testing"
	"time"
)

func togglClient(t *testing.T) *TimeEntryClient {
	tu := &gttest.TestUtil{}
	client := tu.MockClient(t)
	return NewClient(client)
}

func TestTimeEntryDelete(t *testing.T) {
	tClient := togglClient(t)
	err := tClient.Delete(1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTimeEntryList(t *testing.T) {
	tClient := togglClient(t)
	te, err := tClient.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(te) < 1 {
		t.Fatal("<1")
	}

}

func TestTimeEntryCreate(t *testing.T) {
	tClient := togglClient(t)

	te := &TimeEntry{}
	te.Billable = false
	te.Duration = 1200
	te.Pid = 123
	te.Wid = 777
	te.Description = "Meeting with possible clients"
	te.Tags = []string{"billed"}

	nTe, err := tClient.Create(te)

	if err != nil {
		t.Fatal(err)
	}
	if nTe.Id != 3 {
		t.Error("!= 3")
	}
}

func TestTimeEntryUpdate(t *testing.T) {
	tClient := togglClient(t)
	te, err := tClient.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	te.Description = "new"
	nTe, err := tClient.Update(te)
	if err != nil {
		t.Fatal(err)
	}
	if nTe.Description != "new" {
		t.Error("!= new")
	}
}

func TestTimeEntryGet(t *testing.T) {
	tClient := togglClient(t)

	timeentry, err := tClient.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if timeentry.Id != 1 {
		t.Error("!= 1")
	}

	st, err := time.Parse(time.RFC3339, "2013-02-27T01:24:00+00:00")

	if err != nil {
		t.Fatal(err)
	}
	diff := st.Sub(timeentry.Start)
	if diff != 0 {
		t.Errorf("!= %s", diff)
	}
	st, err = time.Parse(time.RFC3339, "2013-02-27T07:24:00+00:00")
	diff = st.Sub(timeentry.Stop)
	if diff != 0 {
		t.Errorf("!= %s", diff)
	}
}

/* Disabled because doesn't work on Windoze
func TestTimeEntryGetTimeRange(t *testing.T) {
	tClient := togglClient(t)

	start, _ := time.Parse(time.RFC3339, "2013-03-10T15:42:46+02:00")
	end, _ := time.Parse(time.RFC3339, "2013-03-12T15:42:46+02:00")

	te, err := tClient.GetRange(start, end)
	if err != nil {
		t.Fatal(err)
	}
	if len(te) < 1 {
		t.Fatal("<1")
	}
	timeentry := te[0]
	if timeentry.Id != 4 {
		t.Error("!= 4")
	}
}
*/

func BenchmarkTimeEntryClient_Get(b *testing.B) {
	b.ReportAllocs()
	tClient := togglClient(nil)

	for i := 0; i < b.N; i++ {
		if _, err := tClient.Get(1); err != nil {
			b.Fatalf("Get: %v", err)
		}
	}
}
