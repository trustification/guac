//
// Copyright 2023 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package guesser

import (
	"github.com/guacsec/guac/pkg/handler/processor"
	guaccasf "github.com/guacsec/guac/pkg/handler/processor/csaf"
	"github.com/openvex/go-vex/pkg/csaf"
)

type csafRedHatTypeGuesser struct{}

func (_ *csafRedHatTypeGuesser) GuessDocumentType(blob []byte, format processor.FormatType) processor.DocumentType {
	switch format {
	case processor.FormatJSON:
		// Decode the BOM
		var decoded csaf.CSAF
		err := guaccasf.UnmarshallFixingDates(blob, &decoded)
		// TODO fix the check once decoded.Document.Publisher will be available
		if err == nil && decoded.Document.Tracking.ID != "" /*&& decoded.Document.Publisher.Namespace*/ {
			return processor.DocumentCsaf
		}
	}
	return processor.DocumentUnknown
}
