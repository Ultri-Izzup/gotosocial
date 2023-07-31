// Code generated by astool. DO NOT EDIT.

package typehashtag

import (
	"fmt"
	vocab "github.com/superseriousbusiness/activity/streams/vocab"
	"strings"
)

//
//
//   {
//     "content": "example",
//     "id": "https://example.com/@alice/hello-world",
//     "tag": [
//       {
//         "href": "https://example.com/hashtag/example",
//         "name": "#example",
//         "type": "Hashtag"
//       }
//     ],
//     "type": "Note"
//   }
type TootHashtag struct {
	ActivityStreamsAttributedTo vocab.ActivityStreamsAttributedToProperty
	ActivityStreamsHeight       vocab.ActivityStreamsHeightProperty
	ActivityStreamsHref         vocab.ActivityStreamsHrefProperty
	ActivityStreamsHreflang     vocab.ActivityStreamsHreflangProperty
	JSONLDId                    vocab.JSONLDIdProperty
	ActivityStreamsMediaType    vocab.ActivityStreamsMediaTypeProperty
	ActivityStreamsName         vocab.ActivityStreamsNameProperty
	ActivityStreamsPreview      vocab.ActivityStreamsPreviewProperty
	ActivityStreamsRel          vocab.ActivityStreamsRelProperty
	ActivityStreamsSummary      vocab.ActivityStreamsSummaryProperty
	JSONLDType                  vocab.JSONLDTypeProperty
	ActivityStreamsWidth        vocab.ActivityStreamsWidthProperty
	alias                       string
	unknown                     map[string]interface{}
}

// DeserializeHashtag creates a Hashtag from a map representation that has been
// unmarshalled from a text or binary format.
func DeserializeHashtag(m map[string]interface{}, aliasMap map[string]string) (*TootHashtag, error) {
	alias := ""
	aliasPrefix := ""
	if a, ok := aliasMap["http://joinmastodon.org/ns"]; ok {
		alias = a
		aliasPrefix = a + ":"
	}
	this := &TootHashtag{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}
	if typeValue, ok := m["type"]; !ok {
		return nil, fmt.Errorf("no \"type\" property in map")
	} else if typeString, ok := typeValue.(string); ok {
		typeName := strings.TrimPrefix(typeString, aliasPrefix)
		if typeName != "Hashtag" {
			return nil, fmt.Errorf("\"type\" property is not of %q type: %s", "Hashtag", typeName)
		}
		// Fall through, success in finding a proper Type
	} else if arrType, ok := typeValue.([]interface{}); ok {
		found := false
		for _, elemVal := range arrType {
			if typeString, ok := elemVal.(string); ok && strings.TrimPrefix(typeString, aliasPrefix) == "Hashtag" {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("could not find a \"type\" property of value %q", "Hashtag")
		}
		// Fall through, success in finding a proper Type
	} else {
		return nil, fmt.Errorf("\"type\" property is unrecognized type: %T", typeValue)
	}
	// Begin: Known property deserialization
	if p, err := mgr.DeserializeAttributedToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsAttributedTo = p
	}
	if p, err := mgr.DeserializeHeightPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsHeight = p
	}
	if p, err := mgr.DeserializeHrefPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsHref = p
	}
	if p, err := mgr.DeserializeHreflangPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsHreflang = p
	}
	if p, err := mgr.DeserializeIdPropertyJSONLD()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.JSONLDId = p
	}
	if p, err := mgr.DeserializeMediaTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsMediaType = p
	}
	if p, err := mgr.DeserializeNamePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsName = p
	}
	if p, err := mgr.DeserializePreviewPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsPreview = p
	}
	if p, err := mgr.DeserializeRelPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsRel = p
	}
	if p, err := mgr.DeserializeSummaryPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsSummary = p
	}
	if p, err := mgr.DeserializeTypePropertyJSONLD()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.JSONLDType = p
	}
	if p, err := mgr.DeserializeWidthPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsWidth = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "attributedTo" {
			continue
		} else if k == "height" {
			continue
		} else if k == "href" {
			continue
		} else if k == "hreflang" {
			continue
		} else if k == "id" {
			continue
		} else if k == "mediaType" {
			continue
		} else if k == "name" {
			continue
		} else if k == "nameMap" {
			continue
		} else if k == "preview" {
			continue
		} else if k == "rel" {
			continue
		} else if k == "summary" {
			continue
		} else if k == "summaryMap" {
			continue
		} else if k == "type" {
			continue
		} else if k == "width" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// HashtagIsDisjointWith returns true if the other provided type is disjoint with
// the Hashtag type.
func HashtagIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Accept", "Activity", "Add", "Announce", "Application", "Arrive", "Article", "Audio", "Block", "Collection", "CollectionPage", "Create", "Delete", "Dislike", "Document", "Emoji", "Event", "Flag", "Follow", "Group", "IdentityProof", "Ignore", "Image", "IntransitiveActivity", "Invite", "Join", "Leave", "Like", "Listen", "Move", "Note", "Object", "Offer", "OrderedCollection", "OrderedCollectionPage", "OrderedCollectionPage", "Organization", "Page", "Person", "Place", "Profile", "PropertyValue", "Question", "Read", "Reject", "Relationship", "Remove", "Service", "TentativeAccept", "TentativeReject", "Tombstone", "Travel", "Undo", "Update", "Video", "View"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetTypeName() {
			return true
		}
	}
	return false
}

// HashtagIsExtendedBy returns true if the other provided type extends from the
// Hashtag type. Note that it returns false if the types are the same; see the
// "IsOrExtendsHashtag" variant instead.
func HashtagIsExtendedBy(other vocab.Type) bool {
	// Shortcut implementation: is not extended by anything.
	return false
}

// IsOrExtendsHashtag returns true if the other provided type is the Hashtag type
// or extends from the Hashtag type.
func IsOrExtendsHashtag(other vocab.Type) bool {
	if other.GetTypeName() == "Hashtag" {
		return true
	}
	return HashtagIsExtendedBy(other)
}

// NewTootHashtag creates a new Hashtag type
func NewTootHashtag() *TootHashtag {
	typeProp := typePropertyConstructor()
	typeProp.AppendXMLSchemaString("Hashtag")
	return &TootHashtag{
		JSONLDType: typeProp,
		alias:      "",
		unknown:    make(map[string]interface{}),
	}
}

// TootHashtagExtends returns true if the Hashtag type extends from the other type.
func TootHashtagExtends(other vocab.Type) bool {
	extensions := []string{"Link"}
	for _, ext := range extensions {
		if ext == other.GetTypeName() {
			return true
		}
	}
	return false
}

// GetActivityStreamsAttributedTo returns the "attributedTo" property if it
// exists, and nil otherwise.
func (this TootHashtag) GetActivityStreamsAttributedTo() vocab.ActivityStreamsAttributedToProperty {
	return this.ActivityStreamsAttributedTo
}

// GetActivityStreamsHeight returns the "height" property if it exists, and nil
// otherwise.
func (this TootHashtag) GetActivityStreamsHeight() vocab.ActivityStreamsHeightProperty {
	return this.ActivityStreamsHeight
}

// GetActivityStreamsHref returns the "href" property if it exists, and nil
// otherwise.
func (this TootHashtag) GetActivityStreamsHref() vocab.ActivityStreamsHrefProperty {
	return this.ActivityStreamsHref
}

// GetActivityStreamsHreflang returns the "hreflang" property if it exists, and
// nil otherwise.
func (this TootHashtag) GetActivityStreamsHreflang() vocab.ActivityStreamsHreflangProperty {
	return this.ActivityStreamsHreflang
}

// GetActivityStreamsMediaType returns the "mediaType" property if it exists, and
// nil otherwise.
func (this TootHashtag) GetActivityStreamsMediaType() vocab.ActivityStreamsMediaTypeProperty {
	return this.ActivityStreamsMediaType
}

// GetActivityStreamsName returns the "name" property if it exists, and nil
// otherwise.
func (this TootHashtag) GetActivityStreamsName() vocab.ActivityStreamsNameProperty {
	return this.ActivityStreamsName
}

// GetActivityStreamsPreview returns the "preview" property if it exists, and nil
// otherwise.
func (this TootHashtag) GetActivityStreamsPreview() vocab.ActivityStreamsPreviewProperty {
	return this.ActivityStreamsPreview
}

// GetActivityStreamsRel returns the "rel" property if it exists, and nil
// otherwise.
func (this TootHashtag) GetActivityStreamsRel() vocab.ActivityStreamsRelProperty {
	return this.ActivityStreamsRel
}

// GetActivityStreamsSummary returns the "summary" property if it exists, and nil
// otherwise.
func (this TootHashtag) GetActivityStreamsSummary() vocab.ActivityStreamsSummaryProperty {
	return this.ActivityStreamsSummary
}

// GetActivityStreamsWidth returns the "width" property if it exists, and nil
// otherwise.
func (this TootHashtag) GetActivityStreamsWidth() vocab.ActivityStreamsWidthProperty {
	return this.ActivityStreamsWidth
}

// GetJSONLDId returns the "id" property if it exists, and nil otherwise.
func (this TootHashtag) GetJSONLDId() vocab.JSONLDIdProperty {
	return this.JSONLDId
}

// GetJSONLDType returns the "type" property if it exists, and nil otherwise.
func (this TootHashtag) GetJSONLDType() vocab.JSONLDTypeProperty {
	return this.JSONLDType
}

// GetTypeName returns the name of this type.
func (this TootHashtag) GetTypeName() string {
	return "Hashtag"
}

// GetUnknownProperties returns the unknown properties for the Hashtag type. Note
// that this should not be used by app developers. It is only used to help
// determine which implementation is LessThan the other. Developers who are
// creating a different implementation of this type's interface can use this
// method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this TootHashtag) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// IsExtending returns true if the Hashtag type extends from the other type.
func (this TootHashtag) IsExtending(other vocab.Type) bool {
	return TootHashtagExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this TootHashtag) JSONLDContext() map[string]string {
	m := map[string]string{"http://joinmastodon.org/ns": this.alias}
	m = this.helperJSONLDContext(this.ActivityStreamsAttributedTo, m)
	m = this.helperJSONLDContext(this.ActivityStreamsHeight, m)
	m = this.helperJSONLDContext(this.ActivityStreamsHref, m)
	m = this.helperJSONLDContext(this.ActivityStreamsHreflang, m)
	m = this.helperJSONLDContext(this.JSONLDId, m)
	m = this.helperJSONLDContext(this.ActivityStreamsMediaType, m)
	m = this.helperJSONLDContext(this.ActivityStreamsName, m)
	m = this.helperJSONLDContext(this.ActivityStreamsPreview, m)
	m = this.helperJSONLDContext(this.ActivityStreamsRel, m)
	m = this.helperJSONLDContext(this.ActivityStreamsSummary, m)
	m = this.helperJSONLDContext(this.JSONLDType, m)
	m = this.helperJSONLDContext(this.ActivityStreamsWidth, m)

	return m
}

// LessThan computes if this Hashtag is lesser, with an arbitrary but stable
// determination.
func (this TootHashtag) LessThan(o vocab.TootHashtag) bool {
	// Begin: Compare known properties
	// Compare property "attributedTo"
	if lhs, rhs := this.ActivityStreamsAttributedTo, o.GetActivityStreamsAttributedTo(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "height"
	if lhs, rhs := this.ActivityStreamsHeight, o.GetActivityStreamsHeight(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "href"
	if lhs, rhs := this.ActivityStreamsHref, o.GetActivityStreamsHref(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "hreflang"
	if lhs, rhs := this.ActivityStreamsHreflang, o.GetActivityStreamsHreflang(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "id"
	if lhs, rhs := this.JSONLDId, o.GetJSONLDId(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "mediaType"
	if lhs, rhs := this.ActivityStreamsMediaType, o.GetActivityStreamsMediaType(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "name"
	if lhs, rhs := this.ActivityStreamsName, o.GetActivityStreamsName(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "preview"
	if lhs, rhs := this.ActivityStreamsPreview, o.GetActivityStreamsPreview(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "rel"
	if lhs, rhs := this.ActivityStreamsRel, o.GetActivityStreamsRel(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "summary"
	if lhs, rhs := this.ActivityStreamsSummary, o.GetActivityStreamsSummary(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "type"
	if lhs, rhs := this.JSONLDType, o.GetJSONLDType(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "width"
	if lhs, rhs := this.ActivityStreamsWidth, o.GetActivityStreamsWidth(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// End: Compare known properties

	// Begin: Compare unknown properties (only by number of them)
	if len(this.unknown) < len(o.GetUnknownProperties()) {
		return true
	} else if len(o.GetUnknownProperties()) < len(this.unknown) {
		return false
	} // End: Compare unknown properties (only by number of them)

	// All properties are the same.
	return false
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format.
func (this TootHashtag) Serialize() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	typeName := "Hashtag"
	if len(this.alias) > 0 {
		typeName = this.alias + ":" + "Hashtag"
	}
	m["type"] = typeName
	// Begin: Serialize known properties
	// Maybe serialize property "attributedTo"
	if this.ActivityStreamsAttributedTo != nil {
		if i, err := this.ActivityStreamsAttributedTo.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsAttributedTo.Name()] = i
		}
	}
	// Maybe serialize property "height"
	if this.ActivityStreamsHeight != nil {
		if i, err := this.ActivityStreamsHeight.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsHeight.Name()] = i
		}
	}
	// Maybe serialize property "href"
	if this.ActivityStreamsHref != nil {
		if i, err := this.ActivityStreamsHref.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsHref.Name()] = i
		}
	}
	// Maybe serialize property "hreflang"
	if this.ActivityStreamsHreflang != nil {
		if i, err := this.ActivityStreamsHreflang.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsHreflang.Name()] = i
		}
	}
	// Maybe serialize property "id"
	if this.JSONLDId != nil {
		if i, err := this.JSONLDId.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.JSONLDId.Name()] = i
		}
	}
	// Maybe serialize property "mediaType"
	if this.ActivityStreamsMediaType != nil {
		if i, err := this.ActivityStreamsMediaType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsMediaType.Name()] = i
		}
	}
	// Maybe serialize property "name"
	if this.ActivityStreamsName != nil {
		if i, err := this.ActivityStreamsName.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsName.Name()] = i
		}
	}
	// Maybe serialize property "preview"
	if this.ActivityStreamsPreview != nil {
		if i, err := this.ActivityStreamsPreview.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsPreview.Name()] = i
		}
	}
	// Maybe serialize property "rel"
	if this.ActivityStreamsRel != nil {
		if i, err := this.ActivityStreamsRel.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsRel.Name()] = i
		}
	}
	// Maybe serialize property "summary"
	if this.ActivityStreamsSummary != nil {
		if i, err := this.ActivityStreamsSummary.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsSummary.Name()] = i
		}
	}
	// Maybe serialize property "type"
	if this.JSONLDType != nil {
		if i, err := this.JSONLDType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.JSONLDType.Name()] = i
		}
	}
	// Maybe serialize property "width"
	if this.ActivityStreamsWidth != nil {
		if i, err := this.ActivityStreamsWidth.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsWidth.Name()] = i
		}
	}
	// End: Serialize known properties

	// Begin: Serialize unknown properties
	for k, v := range this.unknown {
		// To be safe, ensure we aren't overwriting a known property
		if _, has := m[k]; !has {
			m[k] = v
		}
	}
	// End: Serialize unknown properties

	return m, nil
}

// SetActivityStreamsAttributedTo sets the "attributedTo" property.
func (this *TootHashtag) SetActivityStreamsAttributedTo(i vocab.ActivityStreamsAttributedToProperty) {
	this.ActivityStreamsAttributedTo = i
}

// SetActivityStreamsHeight sets the "height" property.
func (this *TootHashtag) SetActivityStreamsHeight(i vocab.ActivityStreamsHeightProperty) {
	this.ActivityStreamsHeight = i
}

// SetActivityStreamsHref sets the "href" property.
func (this *TootHashtag) SetActivityStreamsHref(i vocab.ActivityStreamsHrefProperty) {
	this.ActivityStreamsHref = i
}

// SetActivityStreamsHreflang sets the "hreflang" property.
func (this *TootHashtag) SetActivityStreamsHreflang(i vocab.ActivityStreamsHreflangProperty) {
	this.ActivityStreamsHreflang = i
}

// SetActivityStreamsMediaType sets the "mediaType" property.
func (this *TootHashtag) SetActivityStreamsMediaType(i vocab.ActivityStreamsMediaTypeProperty) {
	this.ActivityStreamsMediaType = i
}

// SetActivityStreamsName sets the "name" property.
func (this *TootHashtag) SetActivityStreamsName(i vocab.ActivityStreamsNameProperty) {
	this.ActivityStreamsName = i
}

// SetActivityStreamsPreview sets the "preview" property.
func (this *TootHashtag) SetActivityStreamsPreview(i vocab.ActivityStreamsPreviewProperty) {
	this.ActivityStreamsPreview = i
}

// SetActivityStreamsRel sets the "rel" property.
func (this *TootHashtag) SetActivityStreamsRel(i vocab.ActivityStreamsRelProperty) {
	this.ActivityStreamsRel = i
}

// SetActivityStreamsSummary sets the "summary" property.
func (this *TootHashtag) SetActivityStreamsSummary(i vocab.ActivityStreamsSummaryProperty) {
	this.ActivityStreamsSummary = i
}

// SetActivityStreamsWidth sets the "width" property.
func (this *TootHashtag) SetActivityStreamsWidth(i vocab.ActivityStreamsWidthProperty) {
	this.ActivityStreamsWidth = i
}

// SetJSONLDId sets the "id" property.
func (this *TootHashtag) SetJSONLDId(i vocab.JSONLDIdProperty) {
	this.JSONLDId = i
}

// SetJSONLDType sets the "type" property.
func (this *TootHashtag) SetJSONLDType(i vocab.JSONLDTypeProperty) {
	this.JSONLDType = i
}

// VocabularyURI returns the vocabulary's URI as a string.
func (this TootHashtag) VocabularyURI() string {
	return "http://joinmastodon.org/ns"
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this TootHashtag) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
	if i == nil {
		return toMerge
	}
	for k, v := range i.JSONLDContext() {
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		toMerge[k] = v
	}
	return toMerge
}