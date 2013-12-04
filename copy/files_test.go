package copy

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var (
	fileService *FileService
)

func setupFileService(t *testing.T) {
	setup(t)
	fileService = NewFileService(client)
}

func tearDownFileService() {
	defer tearDown()
}

// Checks json decoding for the meta object
func TestJsonMetaDecoding(t *testing.T) {
	setupFileService(t)
	defer tearDownFileService()
	mux.HandleFunc("/meta",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w,
				`{
               "id":"\/",
               "path":"\/",
               "name":"Copy",
               "type":"root",
               "stub":false,
               "children":[
                  {
                     "name":"Personal Files",
                     "type":"copy",
                     "id":"\/copy",
                     "path":"\/",
                     "stub":true,
                     "counts":{
                        "new":0,
                        "viewed":0,
                        "hidden":0
                     }
                  }
               ],
               "children_count":1,
               "link_name":"link test",
               "token":"32234dsad",
               "permissions":"all",
               "public":true,
               "size":3123123,
               "date_last_synced":32131232,
               "share":true,
               "recipient_confirmed":true,
               "object_available":true,
               "links": [
                   {
                        "id":"link1",
                        "public":true,
                        "expires":true,
                        "expired":true,
                        "url":"dsafdsfdsaxfwf",
                        "url_short":"dsadsad",
                        "recipients": [
                            {
                                "contact_Type":"gfgdfd",
                                "contact_id":"fgffsd",
                                "contact_source":"htgdffvdb",
                                "user_id":"3343",
                                "first_name":"ffgfgf",
                                "last_name":"grfesa",
                                "email":"fsdfdsfds",
                                "permissions":"all",
                                "emails": [
                                     {
                                            "confirmed":true,
                                            "primary":true,
                                            "email":"thomashunter@example.com",
                                            "gravatar":"eca957c6552e783627a0ced1035e1888"
                                    }
                                ]
                            }
                        ],
                        "creator_id":"htgdffsdd",
                        "confirmation_required": true
                    }
               ],
               "revisions": [
                    {
                        "revision_id":"231312",
                        "modified_time":"32324",
                        "size":31232,
                        "latest":true,
                        "conflict":4324,
                        "id":"dsdsd",
                        "type":"sdsad",
                        "creator":{
                            "user_id":"44342",
                            "created_time":323423,
                            "email":"fdfdsf@dsadsa.com",
                            "first_name":"sadasd",
                            "last_name":"sdsadsafds",
                            "confirmed":true
                        }
                    }
                ],
                "url":"dasdsafdasddfdf",
                "revision_id":31312,
                "thumb":"test thumb",
                "thumb_original_dimensions":{
                    "width":32432,
                    "height":53543
                }
            }`)
		},
	)

	fileMeta, _ := fileService.GetTopLevelMeta()

	perfectFileMeta := Meta{
		Id:   "/",
		Path: "/",
		Name: "Copy",
		Type: "root",
		Stub: false,
		Children: []Meta{
			Meta{
				Id:   "/copy",
				Path: "/",
				Name: "Personal Files",
				Type: "copy",
				Stub: true,
				Counts: Count{
					New:    0,
					Viewed: 0,
					Hidden: 0,
				},
			},
		},
		ChildrenCount:      1,
		LinkName:           "link test",
		Token:              "32234dsad",
		Permissions:        "all",
		Public:             true,
		Size:               3123123,
		DateLastSynced:     32131232,
		Share:              true,
		RecipientConfirmed: true,
		ObjectAvailable:    true,
		Links: []Link{
			Link{
				Id:       "link1",
				Public:   true,
				Expires:  true,
				Expired:  true,
				Url:      "dsafdsfdsaxfwf",
				UrlShort: "dsadsad",
				Recipients: []Recipient{
					Recipient{
						ContactType:   "gfgdfd",
						ContactId:     "fgffsd",
						ContactSource: "htgdffvdb",
						UserId:        "3343",
						FirstName:     "ffgfgf",
						LastName:      "grfesa",
						Email:         "fsdfdsfds",
						Permissions:   "all",
						Emails: []Email{
							Email{
								Confirmed: true,
								Primary:   true,
								Email:     "thomashunter@example.com",
								Gravatar:  "eca957c6552e783627a0ced1035e1888",
							},
						},
					},
				},
				CreatorId:            "htgdffsdd",
				ConfirmationRequired: true,
			},
		},
		Revisions: []Revision{
			Revision{
				RevisionId:   "231312",
				ModifiedTime: "32324",
				Size:         31232,
				Latest:       true,
				Conflict:     4324,
				Id:           "dsdsd",
				Type:         "sdsad",
				Creator: Creator{
					UserId:      "44342",
					CreatedTime: 323423,
					Email:       "fdfdsf@dsadsa.com",
					FirstName:   "sadasd",
					LastName:    "sdsadsafds",
					Confirmed:   true,
				},
			},
		},
		Url:        "dasdsafdasddfdf",
		RevisionId: 31312,
		Thumb:      "test thumb",
		ThumbOriginalDimensions: ThumbOriginalDimensions{
			Width:  32432,
			Height: 53543,
		},
	}

	// Are bouth content equal?
	if !reflect.DeepEqual(*fileMeta, perfectFileMeta) {
		t.Errorf("Metas are not equal")
	}
}

func TestGetMeta(t *testing.T) {
	setupFileService(t)
	defer tearDownFileService()

	mux.HandleFunc("/"+fmt.Sprintf(getMetaSuffix, "testing"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w,
				`{
               "id":"\/copy\/testing",
               "path":"\/testing",
               "name":"testing",
               "type":"dir",
               "size":null,
               "date_last_synced":1386150047,
               "modified_time":1386150047,
               "stub":false,
               "recipient_confirmed":false,
               "counts":[

               ],
               "mime_type":"",
               "link_name":null,
               "token":null,
               "creator_id":null,
               "permissions":null,
               "syncing":false,
               "public":false,
               "object_available":true,
               "links":[

               ],
               "url":"https:\/\/copy.com\/web\/users\/user-8129109\/copy\/testing",
               "thumb":null,
               "share":null,
               "children":[
                  {
                     "id":"\/copy\/testing\/random.txt",
                     "path":"\/testing\/random.txt",
                     "name":"random.txt",
                     "type":"file",
                     "size":1258291200,
                     "date_last_synced":1386151250,
                     "modified_time":1385993169,
                     "stub":true,
                     "recipient_confirmed":false,
                     "counts":[

                     ],
                     "mime_type":"text\/plain",
                     "link_name":null,
                     "token":null,
                     "creator_id":null,
                     "permissions":null,
                     "syncing":false,
                     "public":false,
                     "object_available":true,
                     "links":[

                     ],
                     "url":"https:\/\/copy.com\/web\/users\/user-8129109\/copy\/testing\/random.txt",
                     "revision":32,
                     "thumb":null,
                     "share":null,
                     "list_index":0
                  }
               ],
               "children_count":1
            }`)
		},
	)

	fileMeta, _ := fileService.GetMeta("testing")

	perfectFileMeta := Meta{
		Id:                 "/copy/testing",
		Path:               "/testing",
		Name:               "testing",
		Type:               "dir",
		DateLastSynced:     1386150047,
		ModifiedTime:       1386150047,
		Stub:               false,
		RecipientConfirmed: false,
		Syncing:            false,
		Public:             false,
		ObjectAvailable:    true,
		Url:                "https://copy.com/web/users/user-8129109/copy/testing",
		Links:              []Link{}, // for Deep equal nil and empty slice aren't the same
		Children: []Meta{
			Meta{
				Id:                 "/copy/testing/random.txt",
				Path:               "/testing/random.txt",
				Name:               "random.txt",
				Type:               "file",
				Size:               1258291200,
				DateLastSynced:     1386151250,
				ModifiedTime:       1385993169,
				Stub:               true,
				RecipientConfirmed: false,
				MimeType:           "text/plain",
				Syncing:            false,
				Public:             false,
				ObjectAvailable:    true,
				Url:                "https://copy.com/web/users/user-8129109/copy/testing/random.txt",
				Revision:           32,
				ListIndex:          0,
				Links:              []Link{}, // for Deep equal nil and empty slice aren't the same

			},
		},
		ChildrenCount: 1,
	}

	// Are bouth content equal?
	if !reflect.DeepEqual(*fileMeta, perfectFileMeta) {
		t.Errorf("Metas are not equal")
	}

	// Check error in request

	//Prepare the neccesary data
	/*appToken := os.Getenv("APP_TOKEN")
	appSecret := os.Getenv("APP_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	// Create the client
	client, _ := NewDefaultClient(appToken, appSecret, accessToken, accessSecret)
	fs := NewFileService(client)

	fs.GetMeta("testing/")*/

}

// Checks json decoding for the meta object
func TestGetFile(t *testing.T) {
	setupFileService(t)
	defer tearDownFileService()

	filename := "client_test.go"
	// Read the file to test
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		t.Error(err.Error())
	}

	mux.HandleFunc(strings.Join([]string{"", filesTopLevelSuffix, filename}, "/"),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			w.Write(file)
		},
	)

	fileReader, _ := fileService.GetFile(filename)
	defer fileReader.Close()

	file2, err := ioutil.ReadAll(fileReader)

	if err != nil {
		t.Error(err.Error())
	}

	if !bytes.Equal(file, file2) {
		t.Errorf("contents are not equal")
	}
}

func TestFileUpload(t *testing.T) {
	setupFileService(t)
	defer tearDownFileService()

	filePath := "files_test.go"
	upPath := "tests/uploads"

	// Read the file to test
	origFile, err := ioutil.ReadFile(filePath)

	if err != nil {
		t.Error(err.Error())
	}

	resPath := strings.Join([]string{"", filesTopLevelSuffix, upPath}, "/")

	mux.HandleFunc(resPath,
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			// Check that upload is ok
			r.ParseMultipartForm(100000)
			form := r.MultipartForm

			files, _ := form.File["file"]
			file, _ := files[0].Open()
			defer file.Close()

			buf := new(bytes.Buffer)
			io.Copy(buf, file)

			if !bytes.Equal(origFile, buf.Bytes()) {
				t.Errorf("contents are not equal")
			}
		},
	)

	err = fileService.UploadFile(filePath, strings.Join([]string{upPath, filePath}, "/"), true)

	if err != nil {
		t.Error(err.Error())
	}

}
