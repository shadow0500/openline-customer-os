package dataloader

import (
	"context"
	"errors"
	"github.com/graph-gophers/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"reflect"
	"time"
)

const noteContextTimeout = 10 * time.Second

func (i *Loaders) GetMentionedByNotesForIssue(ctx context.Context, noteId string) (*entity.NoteEntities, error) {
	thunk := i.MentionedByNotesForIssue.Load(ctx, dataloader.StringKey(noteId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	resultObj := result.(entity.NoteEntities)
	return &resultObj, nil
}

func (b *noteBatcher) getMentionedByNotesForIssue(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids, keyOrder := sortKeys(keys)

	ctx, cancel := context.WithTimeout(ctx, noteContextTimeout)
	defer cancel()

	noteEntitiesPtr, err := b.noteService.GetMentionedByNotesForIssues(ctx, ids)
	if err != nil {
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get note entities for issues")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	noteEntitiesByIssueId := make(map[string]entity.NoteEntities)
	for _, val := range *noteEntitiesPtr {
		if list, ok := noteEntitiesByIssueId[val.GetDataloaderKey()]; ok {
			noteEntitiesByIssueId[val.GetDataloaderKey()] = append(list, val)
		} else {
			noteEntitiesByIssueId[val.GetDataloaderKey()] = entity.NoteEntities{val}
		}
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for issueId, record := range noteEntitiesByIssueId {
		if ix, ok := keyOrder[issueId]; ok {
			results[ix] = &dataloader.Result{Data: record, Error: nil}
			delete(keyOrder, issueId)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: entity.NoteEntities{}, Error: nil}
	}

	if err = assertEntitiesType(results, reflect.TypeOf(entity.NoteEntities{})); err != nil {
		return []*dataloader.Result{{nil, err}}
	}

	return results
}

func (i *Loaders) GetNotesForMeeting(ctx context.Context, meetingId string) (*entity.NoteEntities, error) {
	thunk := i.NotesForMeeting.Load(ctx, dataloader.StringKey(meetingId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	resultObj := result.(entity.NoteEntities)
	return &resultObj, nil
}

func (b *noteBatcher) getNotesForMeetings(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids, keyOrder := sortKeys(keys)

	ctx, cancel := context.WithTimeout(ctx, noteContextTimeout)
	defer cancel()

	noteEntitiesPtr, err := b.noteService.GetNotesForMeetings(ctx, ids)
	if err != nil {
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get noted entities for notes")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	notesForMeetings := make(map[string]entity.NoteEntities)
	for _, val := range *noteEntitiesPtr {
		if list, ok := notesForMeetings[val.GetDataloaderKey()]; ok {
			notesForMeetings[val.GetDataloaderKey()] = append(list, val)
		} else {
			notesForMeetings[val.GetDataloaderKey()] = entity.NoteEntities{val}
		}
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for contactId, record := range notesForMeetings {
		if ix, ok := keyOrder[contactId]; ok {
			results[ix] = &dataloader.Result{Data: record, Error: nil}
			delete(keyOrder, contactId)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: entity.NoteEntities{}, Error: nil}
	}

	if err = assertEntitiesType(results, reflect.TypeOf(entity.NoteEntities{})); err != nil {
		return []*dataloader.Result{{nil, err}}
	}

	return results
}
