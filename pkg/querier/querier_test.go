package querier

import (
	"context"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/go-kit/log"
	"github.com/grafana/dskit/ring"
	"github.com/grafana/dskit/ring/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	commonv1 "github.com/grafana/fire/pkg/gen/common/v1"
	ingestv1 "github.com/grafana/fire/pkg/gen/ingester/v1"
	querierv1 "github.com/grafana/fire/pkg/gen/querier/v1"
	"github.com/grafana/fire/pkg/ingester/clientpool"
	"github.com/grafana/fire/pkg/iter"
	firemodel "github.com/grafana/fire/pkg/model"
	"github.com/grafana/fire/pkg/testhelper"
)

func Test_QuerySampleType(t *testing.T) {
	querier, err := New(Config{
		PoolConfig: clientpool.PoolConfig{ClientCleanupPeriod: 1 * time.Millisecond},
	}, testhelper.NewMockRing([]ring.InstanceDesc{
		{Addr: "1"},
		{Addr: "2"},
		{Addr: "3"},
	}, 3), func(addr string) (client.PoolClient, error) {
		q := newFakeQuerier()
		switch addr {
		case "1":
			q.On("ProfileTypes", mock.Anything, mock.Anything).
				Return(connect.NewResponse(&ingestv1.ProfileTypesResponse{
					ProfileTypes: []*commonv1.ProfileType{
						{ID: "foo"},
						{ID: "bar"},
					},
				}), nil)
		case "2":
			q.On("ProfileTypes", mock.Anything, mock.Anything).
				Return(connect.NewResponse(&ingestv1.ProfileTypesResponse{
					ProfileTypes: []*commonv1.ProfileType{
						{ID: "bar"},
						{ID: "buzz"},
					},
				}), nil)
		case "3":
			q.On("ProfileTypes", mock.Anything, mock.Anything).
				Return(connect.NewResponse(&ingestv1.ProfileTypesResponse{
					ProfileTypes: []*commonv1.ProfileType{
						{ID: "buzz"},
						{ID: "foo"},
					},
				}), nil)
		}
		return q, nil
	}, log.NewLogfmtLogger(os.Stdout))

	require.NoError(t, err)
	out, err := querier.ProfileTypes(context.Background(), connect.NewRequest(&querierv1.ProfileTypesRequest{}))
	ids := make([]string, 0, len(out.Msg.ProfileTypes))
	for _, pt := range out.Msg.ProfileTypes {
		ids = append(ids, pt.ID)
	}
	require.NoError(t, err)
	require.Equal(t, []string{"bar", "buzz", "foo"}, ids)
}

func Test_QueryLabelValues(t *testing.T) {
	req := connect.NewRequest(&querierv1.LabelValuesRequest{Name: "foo"})
	querier, err := New(Config{
		PoolConfig: clientpool.PoolConfig{ClientCleanupPeriod: 1 * time.Millisecond},
	}, testhelper.NewMockRing([]ring.InstanceDesc{
		{Addr: "1"},
		{Addr: "2"},
		{Addr: "3"},
	}, 3), func(addr string) (client.PoolClient, error) {
		q := newFakeQuerier()
		switch addr {
		case "1":
			q.On("LabelValues", mock.Anything, mock.Anything).Return(connect.NewResponse(&ingestv1.LabelValuesResponse{Names: []string{"foo", "bar"}}), nil)
		case "2":
			q.On("LabelValues", mock.Anything, mock.Anything).Return(connect.NewResponse(&ingestv1.LabelValuesResponse{Names: []string{"bar", "buzz"}}), nil)
		case "3":
			q.On("LabelValues", mock.Anything, mock.Anything).Return(connect.NewResponse(&ingestv1.LabelValuesResponse{Names: []string{"buzz", "foo"}}), nil)
		}
		return q, nil
	}, log.NewLogfmtLogger(os.Stdout))

	require.NoError(t, err)
	out, err := querier.LabelValues(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, []string{"bar", "buzz", "foo"}, out.Msg.Names)
}

func Test_QueryLabelNames(t *testing.T) {
	req := connect.NewRequest(&querierv1.LabelNamesRequest{})
	querier, err := New(Config{
		PoolConfig: clientpool.PoolConfig{ClientCleanupPeriod: 1 * time.Millisecond},
	}, testhelper.NewMockRing([]ring.InstanceDesc{
		{Addr: "1"},
		{Addr: "2"},
		{Addr: "3"},
	}, 3), func(addr string) (client.PoolClient, error) {
		q := newFakeQuerier()
		switch addr {
		case "1":
			q.On("LabelNames", mock.Anything, mock.Anything).Return(connect.NewResponse(&ingestv1.LabelNamesResponse{Names: []string{"foo", "bar"}}), nil)
		case "2":
			q.On("LabelNames", mock.Anything, mock.Anything).Return(connect.NewResponse(&ingestv1.LabelNamesResponse{Names: []string{"bar", "buzz"}}), nil)
		case "3":
			q.On("LabelNames", mock.Anything, mock.Anything).Return(connect.NewResponse(&ingestv1.LabelNamesResponse{Names: []string{"buzz", "foo"}}), nil)
		}
		return q, nil
	}, log.NewLogfmtLogger(os.Stdout))

	require.NoError(t, err)
	out, err := querier.LabelNames(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, []string{"bar", "buzz", "foo"}, out.Msg.Names)
}

func Test_Series(t *testing.T) {
	foobarlabels := firemodel.NewLabelsBuilder(nil).Set("foo", "bar")
	foobuzzlabels := firemodel.NewLabelsBuilder(nil).Set("foo", "buzz")
	req := connect.NewRequest(&querierv1.SeriesRequest{Matchers: []string{`{foo="bar"}`}})
	ingesterReponse := connect.NewResponse(&ingestv1.SeriesResponse{LabelsSet: []*commonv1.Labels{
		{Labels: foobarlabels.Labels()},
		{Labels: foobuzzlabels.Labels()},
	}})
	querier, err := New(Config{
		PoolConfig: clientpool.PoolConfig{ClientCleanupPeriod: 1 * time.Millisecond},
	}, testhelper.NewMockRing([]ring.InstanceDesc{
		{Addr: "1"},
		{Addr: "2"},
		{Addr: "3"},
	}, 3), func(addr string) (client.PoolClient, error) {
		q := newFakeQuerier()
		switch addr {
		case "1":
			q.On("Series", mock.Anything, mock.Anything).Return(ingesterReponse, nil)
		case "2":
			q.On("Series", mock.Anything, mock.Anything).Return(ingesterReponse, nil)
		case "3":
			q.On("Series", mock.Anything, mock.Anything).Return(ingesterReponse, nil)
		}
		return q, nil
	}, log.NewLogfmtLogger(os.Stdout))

	require.NoError(t, err)
	out, err := querier.Series(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, []*commonv1.Labels{
		{Labels: foobarlabels.Labels()},
		{Labels: foobuzzlabels.Labels()},
	}, out.Msg.LabelsSet)
}

func Test_SelectMergeStacktraces(t *testing.T) {
	req := connect.NewRequest(&querierv1.SelectMergeStacktracesRequest{
		LabelSelector: `{app="foo"}`,
		ProfileTypeID: "memory:inuse_space:bytes:space:byte",
		Start:         0,
		End:           2,
	})
	bidi1 := newFakeBidiClientStacktraces([]*ingestv1.ProfileSets{
		{
			LabelsSets: []*commonv1.Labels{
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "foo"}},
				},
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "bar"}},
				},
			},
			Profiles: []*ingestv1.SeriesProfile{
				{Timestamp: 1, LabelIndex: 0},
				{Timestamp: 2, LabelIndex: 1},
				{Timestamp: 2, LabelIndex: 0},
			},
		},
	})
	bidi2 := newFakeBidiClientStacktraces([]*ingestv1.ProfileSets{
		{
			LabelsSets: []*commonv1.Labels{
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "foo"}},
				},
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "bar"}},
				},
			},
			Profiles: []*ingestv1.SeriesProfile{
				{Timestamp: 1, LabelIndex: 1},
				{Timestamp: 1, LabelIndex: 0},
				{Timestamp: 2, LabelIndex: 1},
			},
		},
	})
	bidi3 := newFakeBidiClientStacktraces([]*ingestv1.ProfileSets{
		{
			LabelsSets: []*commonv1.Labels{
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "foo"}},
				},
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "bar"}},
				},
			},
			Profiles: []*ingestv1.SeriesProfile{
				{Timestamp: 1, LabelIndex: 1},
				{Timestamp: 1, LabelIndex: 0},
				{Timestamp: 2, LabelIndex: 0},
			},
		},
	})
	querier, err := New(Config{
		PoolConfig: clientpool.PoolConfig{ClientCleanupPeriod: 1 * time.Millisecond},
	}, testhelper.NewMockRing([]ring.InstanceDesc{
		{Addr: "1"},
		{Addr: "2"},
		{Addr: "3"},
	}, 3), func(addr string) (client.PoolClient, error) {
		q := newFakeQuerier()
		switch addr {
		case "1":
			q.On("MergeProfilesStacktraces", mock.Anything).Once().Return(bidi1)
		case "2":
			q.On("MergeProfilesStacktraces", mock.Anything).Once().Return(bidi2)
		case "3":
			q.On("MergeProfilesStacktraces", mock.Anything).Once().Return(bidi3)
		}
		return q, nil
	}, log.NewLogfmtLogger(os.Stdout))
	require.NoError(t, err)
	flame, err := querier.SelectMergeStacktraces(context.Background(), req)
	require.NoError(t, err)

	sort.Strings(flame.Msg.Flamegraph.Names)
	require.Equal(t, []string{"bar", "buzz", "foo", "total"}, flame.Msg.Flamegraph.Names)
	require.Equal(t, []int64{0, 2, 0, 0}, flame.Msg.Flamegraph.Levels[0].Values)
	require.Equal(t, int64(2), flame.Msg.Flamegraph.Total)
	require.Equal(t, int64(2), flame.Msg.Flamegraph.MaxSelf)
	var selected []testProfile
	selected = append(selected, bidi1.kept...)
	selected = append(selected, bidi2.kept...)
	selected = append(selected, bidi3.kept...)
	sort.Slice(selected, func(i, j int) bool {
		if selected[i].Ts == selected[j].Ts {
			return firemodel.CompareLabelPairs(selected[i].Labels.Labels, selected[j].Labels.Labels) < 0
		}
		return selected[i].Ts < selected[j].Ts
	})
	require.Len(t, selected, 4)
	require.Equal(t,
		[]testProfile{
			{Ts: 1, Labels: &commonv1.Labels{Labels: []*commonv1.LabelPair{{Name: "app", Value: "bar"}}}},
			{Ts: 1, Labels: &commonv1.Labels{Labels: []*commonv1.LabelPair{{Name: "app", Value: "foo"}}}},
			{Ts: 2, Labels: &commonv1.Labels{Labels: []*commonv1.LabelPair{{Name: "app", Value: "bar"}}}},
			{Ts: 2, Labels: &commonv1.Labels{Labels: []*commonv1.LabelPair{{Name: "app", Value: "foo"}}}},
		}, selected)
}

func TestSelectSeries(t *testing.T) {
	req := connect.NewRequest(&querierv1.SelectSeriesRequest{
		LabelSelector: `{app="foo"}`,
		ProfileTypeID: "memory:inuse_space:bytes:space:byte",
		Start:         0,
		End:           2,
		Step:          0.001,
	})
	bidi1 := newFakeBidiClientSeries([]*ingestv1.ProfileSets{
		{
			LabelsSets: []*commonv1.Labels{
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "foo"}},
				},
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "bar"}},
				},
			},
			Profiles: []*ingestv1.SeriesProfile{
				{Timestamp: 1, LabelIndex: 0},
				{Timestamp: 2, LabelIndex: 1},
				{Timestamp: 2, LabelIndex: 0},
			},
		},
	}, &commonv1.Series{Labels: foobarlabels, Points: []*commonv1.Point{{V: 1, T: 1}, {V: 2, T: 2}}})
	bidi2 := newFakeBidiClientSeries([]*ingestv1.ProfileSets{
		{
			LabelsSets: []*commonv1.Labels{
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "foo"}},
				},
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "bar"}},
				},
			},
			Profiles: []*ingestv1.SeriesProfile{
				{Timestamp: 1, LabelIndex: 1},
				{Timestamp: 1, LabelIndex: 0},
				{Timestamp: 2, LabelIndex: 1},
			},
		},
	}, &commonv1.Series{Labels: foobarlabels, Points: []*commonv1.Point{{V: 1, T: 1}, {V: 2, T: 2}}})
	bidi3 := newFakeBidiClientSeries([]*ingestv1.ProfileSets{
		{
			LabelsSets: []*commonv1.Labels{
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "foo"}},
				},
				{
					Labels: []*commonv1.LabelPair{{Name: "app", Value: "bar"}},
				},
			},
			Profiles: []*ingestv1.SeriesProfile{
				{Timestamp: 1, LabelIndex: 1},
				{Timestamp: 1, LabelIndex: 0},
				{Timestamp: 2, LabelIndex: 0},
			},
		},
	}, &commonv1.Series{Labels: foobarlabels, Points: []*commonv1.Point{{V: 1, T: 1}, {V: 2, T: 2}}})
	querier, err := New(Config{
		PoolConfig: clientpool.PoolConfig{ClientCleanupPeriod: 1 * time.Millisecond},
	}, testhelper.NewMockRing([]ring.InstanceDesc{
		{Addr: "1"},
		{Addr: "2"},
		{Addr: "3"},
	}, 3), func(addr string) (client.PoolClient, error) {
		q := newFakeQuerier()
		switch addr {
		case "1":
			q.On("MergeProfilesLabels", mock.Anything).Once().Return(bidi1)
		case "2":
			q.On("MergeProfilesLabels", mock.Anything).Once().Return(bidi2)
		case "3":
			q.On("MergeProfilesLabels", mock.Anything).Once().Return(bidi3)
		}
		return q, nil
	}, log.NewLogfmtLogger(os.Stdout))
	require.NoError(t, err)
	res, err := querier.SelectSeries(context.Background(), req)
	require.NoError(t, err)
	// Only 2 results are used since the 3rd not required because of replication.
	testhelper.EqualProto(t, []*commonv1.Series{
		{Labels: foobarlabels, Points: []*commonv1.Point{{V: 2, T: 1}, {V: 4, T: 2}}},
	}, res.Msg.Series)
	var selected []testProfile
	selected = append(selected, bidi1.kept...)
	selected = append(selected, bidi2.kept...)
	selected = append(selected, bidi3.kept...)
	sort.Slice(selected, func(i, j int) bool {
		if selected[i].Ts == selected[j].Ts {
			return firemodel.CompareLabelPairs(selected[i].Labels.Labels, selected[j].Labels.Labels) < 0
		}
		return selected[i].Ts < selected[j].Ts
	})
	require.Len(t, selected, 4)
	require.Equal(t,
		[]testProfile{
			{Ts: 1, Labels: &commonv1.Labels{Labels: []*commonv1.LabelPair{{Name: "app", Value: "bar"}}}},
			{Ts: 1, Labels: &commonv1.Labels{Labels: []*commonv1.LabelPair{{Name: "app", Value: "foo"}}}},
			{Ts: 2, Labels: &commonv1.Labels{Labels: []*commonv1.LabelPair{{Name: "app", Value: "bar"}}}},
			{Ts: 2, Labels: &commonv1.Labels{Labels: []*commonv1.LabelPair{{Name: "app", Value: "foo"}}}},
		}, selected)
}

type fakeQuerierIngester struct {
	mock.Mock
	testhelper.FakePoolClient
}

func newFakeQuerier() *fakeQuerierIngester {
	return &fakeQuerierIngester{}
}

func (f *fakeQuerierIngester) LabelValues(ctx context.Context, req *connect.Request[ingestv1.LabelValuesRequest]) (*connect.Response[ingestv1.LabelValuesResponse], error) {
	var (
		args = f.Called(ctx, req)
		res  *connect.Response[ingestv1.LabelValuesResponse]
		err  error
	)
	if args[0] != nil {
		res = args[0].(*connect.Response[ingestv1.LabelValuesResponse])
	}
	if args[1] != nil {
		err = args.Get(1).(error)
	}
	return res, err
}

func (f *fakeQuerierIngester) LabelNames(ctx context.Context, req *connect.Request[ingestv1.LabelNamesRequest]) (*connect.Response[ingestv1.LabelNamesResponse], error) {
	var (
		args = f.Called(ctx, req)
		res  *connect.Response[ingestv1.LabelNamesResponse]
		err  error
	)
	if args[0] != nil {
		res = args[0].(*connect.Response[ingestv1.LabelNamesResponse])
	}
	if args[1] != nil {
		err = args.Get(1).(error)
	}
	return res, err
}

func (f *fakeQuerierIngester) ProfileTypes(ctx context.Context, req *connect.Request[ingestv1.ProfileTypesRequest]) (*connect.Response[ingestv1.ProfileTypesResponse], error) {
	var (
		args = f.Called(ctx, req)
		res  *connect.Response[ingestv1.ProfileTypesResponse]
		err  error
	)
	if args[0] != nil {
		res = args[0].(*connect.Response[ingestv1.ProfileTypesResponse])
	}
	if args[1] != nil {
		err = args.Get(1).(error)
	}

	return res, err
}

func (f *fakeQuerierIngester) Series(ctx context.Context, req *connect.Request[ingestv1.SeriesRequest]) (*connect.Response[ingestv1.SeriesResponse], error) {
	var (
		args = f.Called(ctx, req)
		res  *connect.Response[ingestv1.SeriesResponse]
		err  error
	)
	if args[0] != nil {
		res = args[0].(*connect.Response[ingestv1.SeriesResponse])
	}
	if args[1] != nil {
		err = args.Get(1).(error)
	}

	return res, err
}

type testProfile struct {
	Ts     int64
	Labels *commonv1.Labels
}

type fakeBidiClientStacktraces struct {
	profiles chan *ingestv1.ProfileSets
	batches  []*ingestv1.ProfileSets
	kept     []testProfile
	cur      *ingestv1.ProfileSets
}

func newFakeBidiClientStacktraces(batches []*ingestv1.ProfileSets) *fakeBidiClientStacktraces {
	res := &fakeBidiClientStacktraces{
		profiles: make(chan *ingestv1.ProfileSets, 1),
	}
	res.profiles <- batches[0]
	batches = batches[1:]
	res.batches = batches
	return res
}

func (f *fakeBidiClientStacktraces) Send(in *ingestv1.MergeProfilesStacktracesRequest) error {
	if in.Request != nil {
		return nil
	}
	for i, b := range in.Profiles {
		if b {
			f.kept = append(f.kept, testProfile{
				Ts:     f.cur.Profiles[i].Timestamp,
				Labels: f.cur.LabelsSets[f.cur.Profiles[i].LabelIndex],
			})
		}
	}
	if len(f.batches) == 0 {
		close(f.profiles)
		return nil
	}
	f.profiles <- f.batches[0]
	f.batches = f.batches[1:]
	return nil
}

func (f *fakeBidiClientStacktraces) Receive() (*ingestv1.MergeProfilesStacktracesResponse, error) {
	profiles := <-f.profiles
	if profiles == nil {
		return &ingestv1.MergeProfilesStacktracesResponse{
			Result: &ingestv1.MergeProfilesStacktracesResult{
				Stacktraces: []*ingestv1.StacktraceSample{
					{FunctionIds: []int32{0, 1, 2}, Value: 1},
				},
				FunctionNames: []string{"foo", "bar", "buzz"},
			},
		}, nil
	}
	f.cur = profiles
	return &ingestv1.MergeProfilesStacktracesResponse{
		SelectedProfiles: profiles,
	}, nil
}
func (f *fakeBidiClientStacktraces) CloseRequest() error  { return nil }
func (f *fakeBidiClientStacktraces) CloseResponse() error { return nil }

type fakeBidiClientSeries struct {
	profiles chan *ingestv1.ProfileSets
	batches  []*ingestv1.ProfileSets
	kept     []testProfile
	cur      *ingestv1.ProfileSets

	result []*commonv1.Series
}

func newFakeBidiClientSeries(batches []*ingestv1.ProfileSets, result ...*commonv1.Series) *fakeBidiClientSeries {
	res := &fakeBidiClientSeries{
		profiles: make(chan *ingestv1.ProfileSets, 1),
	}
	res.profiles <- batches[0]
	batches = batches[1:]
	res.batches = batches
	res.result = result
	return res
}

func (f *fakeBidiClientSeries) Send(in *ingestv1.MergeProfilesLabelsRequest) error {
	if in.Request != nil {
		return nil
	}
	for i, b := range in.Profiles {
		if b {
			f.kept = append(f.kept, testProfile{
				Ts:     f.cur.Profiles[i].Timestamp,
				Labels: f.cur.LabelsSets[f.cur.Profiles[i].LabelIndex],
			})
		}
	}
	if len(f.batches) == 0 {
		close(f.profiles)
		return nil
	}
	f.profiles <- f.batches[0]
	f.batches = f.batches[1:]
	return nil
}

func (f *fakeBidiClientSeries) Receive() (*ingestv1.MergeProfilesLabelsResponse, error) {
	profiles := <-f.profiles
	if profiles == nil {
		return &ingestv1.MergeProfilesLabelsResponse{
			Series: f.result,
		}, nil
	}
	f.cur = profiles
	return &ingestv1.MergeProfilesLabelsResponse{
		SelectedProfiles: profiles,
	}, nil
}
func (f *fakeBidiClientSeries) CloseRequest() error  { return nil }
func (f *fakeBidiClientSeries) CloseResponse() error { return nil }

func (f *fakeQuerierIngester) MergeProfilesStacktraces(ctx context.Context) clientpool.BidiClientMergeProfilesStacktraces {
	var (
		args = f.Called(ctx)
		res  clientpool.BidiClientMergeProfilesStacktraces
	)
	if args[0] != nil {
		res = args[0].(clientpool.BidiClientMergeProfilesStacktraces)
	}

	return res
}

func (f *fakeQuerierIngester) MergeProfilesLabels(ctx context.Context) clientpool.BidiClientMergeProfilesLabels {
	var (
		args = f.Called(ctx)
		res  clientpool.BidiClientMergeProfilesLabels
	)
	if args[0] != nil {
		res = args[0].(clientpool.BidiClientMergeProfilesLabels)
	}

	return res
}

func TestRangeSeries(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   []ProfileValue
		out  []*commonv1.Series
	}{
		{
			name: "single series",
			in: []ProfileValue{
				{Ts: 1, Value: 1},
				{Ts: 1, Value: 1},
				{Ts: 2, Value: 2},
				{Ts: 3, Value: 3},
				{Ts: 4, Value: 4},
				{Ts: 5, Value: 5},
			},
			out: []*commonv1.Series{
				{
					Points: []*commonv1.Point{
						{T: 1, V: 2},
						{T: 2, V: 2},
						{T: 3, V: 3},
						{T: 4, V: 4},
						{T: 5, V: 5},
					},
				},
			},
		},
		{
			name: "multiple series",
			in: []ProfileValue{
				{Ts: 1, Value: 1, Lbs: foobarlabels, LabelsHash: foobarlabels.Hash()},
				{Ts: 1, Value: 1, Lbs: foobuzzlabels, LabelsHash: foobuzzlabels.Hash()},
				{Ts: 2, Value: 1, Lbs: foobarlabels, LabelsHash: foobarlabels.Hash()},
				{Ts: 3, Value: 1, Lbs: foobuzzlabels, LabelsHash: foobuzzlabels.Hash()},
				{Ts: 3, Value: 1, Lbs: foobuzzlabels, LabelsHash: foobuzzlabels.Hash()},
				{Ts: 4, Value: 4, Lbs: foobuzzlabels, LabelsHash: foobuzzlabels.Hash()},
				{Ts: 4, Value: 4, Lbs: foobuzzlabels, LabelsHash: foobuzzlabels.Hash()},
				{Ts: 4, Value: 4, Lbs: foobarlabels, LabelsHash: foobarlabels.Hash()},
				{Ts: 5, Value: 5, Lbs: foobarlabels, LabelsHash: foobarlabels.Hash()},
			},
			out: []*commonv1.Series{
				{
					Labels: foobarlabels,
					Points: []*commonv1.Point{
						{T: 1, V: 1},
						{T: 2, V: 1},
						{T: 4, V: 4},
						{T: 5, V: 5},
					},
				},
				{
					Labels: foobuzzlabels,
					Points: []*commonv1.Point{
						{T: 1, V: 1},
						{T: 3, V: 2},
						{T: 4, V: 8},
					},
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			in := iter.NewSliceIterator(tc.in)
			out := rangeSeries(in, 1, 5, 1)
			testhelper.EqualProto(t, tc.out, out)
		})
	}
}

// The code below can be useful for testing deduping directly to a cluster.
// func TestDedupeLive(t *testing.T) {
// 	clients, err := createClients(context.Background())
// 	require.NoError(t, err)
// 	st, err := dedupe(context.Background(), clients)
// 	require.NoError(t, err)
// 	require.Equal(t, 2, len(st))
// }

// func createClients(ctx context.Context) ([]responseFromIngesters[BidiClientMergeProfilesStacktraces], error) {
// 	var clients []responseFromIngesters[BidiClientMergeProfilesStacktraces]
// 	for i := 1; i < 6; i++ {
// 		addr := fmt.Sprintf("localhost:4%d00", i)
// 		c, err := clientpool.PoolFactory(addr)
// 		if err != nil {
// 			return nil, err
// 		}
// 		res, err := c.Check(ctx, &grpc_health_v1.HealthCheckRequest{
// 			Service: ingestv1.IngesterService_ServiceDesc.ServiceName,
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		if res.Status != grpc_health_v1.HealthCheckResponse_SERVING {
// 			return nil, fmt.Errorf("ingester %s is not serving", addr)
// 		}
// 		bidi := c.(IngesterQueryClient).MergeProfilesStacktraces(ctx)
// 		profileType, err := firemodel.ParseProfileTypeSelector("process_cpu:cpu:nanoseconds:cpu:nanoseconds")
// 		if err != nil {
// 			return nil, err
// 		}
// 		now := time.Now()
// 		err = bidi.Send(&ingestv1.MergeProfilesStacktracesRequest{
// 			Request: &ingestv1.SelectProfilesRequest{
// 				LabelSelector: `{namespace="fire-dev-001"}`,
// 				Type:          profileType,
// 				Start:         int64(model.TimeFromUnixNano(now.Add(-30 * time.Minute).UnixNano())),
// 				End:           int64(model.TimeFromUnixNano(now.UnixNano())),
// 			},
// 		})
// 		if err != nil {
// 			return nil, err
// 		}
// 		clients = append(clients, responseFromIngesters[BidiClientMergeProfilesStacktraces]{
// 			response: bidi,
// 			addr:     addr,
// 		})
// 	}
// 	return clients, nil
// }
