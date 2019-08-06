// DO NOT EDIT - Code generated by firemodel (dev).

import Foundation
import Pring

// [unchanged] Enums provide type safety around string enumerations. Enums are stored in firestore as capitalized strings.
@objc enum Audience: Int {
    case global
    case friends
}

extension Audience: CustomDebugStringConvertible {
    init?(firestoreValue value: Any?) {
        guard let value = value as? String else {
            return nil
        }
        switch value {
        case "GLOBAL":
            self = .global
        case "FRIENDS":
            self = .friends
        default:
            return nil
        }
    }

    var firestoreValue: String? {
        switch self {
        case .global:
            return "GLOBAL"
        case .friends:
            return "FRIENDS"
        }
    }

    var debugDescription: String { return firestoreValue ?? "<INVALID>" }
}
// [new] Enums values may now optionally include associated values. Associatedvalues must have a struct type. Associated values are stored in firestoreunder a period-delimited key, prefixed with the enum field name. The enumcase is always written, even when there is also an associated value. Keysfor enum values other than the active one are not written to firestore.
@objc enum AttachmentContent: Int {
    // e.g. for an Attachment, written as `content = "PLACEHOLDER"`
    case placeholder
    // e.g. for an Attachment, written as `content = "IMOJI", content.imoji = #BYTES#`
    case imoji
    // e.g. for an Attachment, written as `content = "IMOJI", content.imoji = #BYTES#`
    case gram
    // e.g. for an Attachment, written as `content = "UPLOAD", content.upload.token = "oijasdf", content.upload.etc = the_rest`
    case upload
}

extension AttachmentContent: CustomDebugStringConvertible {
    init?(firestoreValue value: Any?) {
        guard let value = value as? String else {
            return nil
        }
        switch value {
        case "PLACEHOLDER":
            self = .placeholder
        case "IMOJI":
            self = .imoji
        case "GRAM":
            self = .gram
        case "UPLOAD":
            self = .upload
        default:
            return nil
        }
    }

    var firestoreValue: String? {
        switch self {
        case .placeholder:
            return "PLACEHOLDER"
        case .imoji:
            return "IMOJI"
        case .gram:
            return "GRAM"
        case .upload:
            return "UPLOAD"
        }
    }

    var debugDescription: String { return firestoreValue ?? "<INVALID>" }
}
@objc enum MessageContent: Int {
    case text
    case photo
}

extension MessageContent: CustomDebugStringConvertible {
    init?(firestoreValue value: Any?) {
        guard let value = value as? String else {
            return nil
        }
        switch value {
        case "TEXT":
            self = .text
        case "PHOTO":
            self = .photo
        default:
            return nil
        }
    }

    var firestoreValue: String? {
        switch self {
        case .text:
            return "TEXT"
        case .photo:
            return "PHOTO"
        }
    }

    var debugDescription: String { return firestoreValue ?? "<INVALID>" }
}

// [unchanged] Structs are used for sharing structures of nested content, whichare stored as Maps in firestore and also accessible via FieldPaths in mostfirestore client libraries.
@objcMembers class Avatar: Pring.Object {
    var url: URL?
    var color: String?
}

@objcMembers class SendMessageRequest: Pring.Object {
    var to: Pring.Reference<Friend> = .init()
    var content: MessageContent?

    override func encode(_ key: String, value: Any?) -> Any? {
        switch key {
        case "content":
            return self.content?.firestoreValue
        default:
            break
        }
        return nil
    }

    override func decode(_ key: String, value: Any?) -> Bool {
        switch key {
        case "content":
            self.content = MessageContent(firestoreValue: value)
        default:
            break
        }
        return false
    }
}

@objcMembers class Upload: Pring.Object {
    var url: URL?
    var mimeType: String?
}

@objcMembers class TextMessageContent: Pring.Object {
    var message: String?
}

@objcMembers class PhotoMessageContent: Pring.Object {
    var caption: String?
    var url: URL?
}

// Models define structures for firestore documents models. Models cannot be used as field types.[new] Models can implement zero or more interfaces.[new] Models can have nested models, which designate nested collections.[new] Models use an inflector to determine their pluralized collection name in firestore.[new] It is now assumed that all fields of all types (other than enums associated values) are optional. In order to facilitate queries on missing fields, missing fields are saved in firestore as explicit null values.
@objcMembers class User: Pring.Object {
    dynamic var username: String?
    dynamic var displayName: String?
    dynamic var avatar: Avatar?

    override func encode(_ key: String, value: Any?) -> Any? {
        switch key {
        case "avatar":
            return self.avatar?.rawValue
        default:
            break
        }
        return nil
    }

    override func decode(_ key: String, value: Any?) -> Bool {
        switch key {
        case "avatar":
          if let value = value as? [String: Any] {
            self.avatar = Avatar(id: "\(0)", value: value)
            return true
          }
        default:
            break
        }
        return false
    }
}

@objcMembers class Gram: Pring.Object {
    dynamic var sharedWith: Audience?
    dynamic var photoUrl: URL?
    dynamic var description: String?

    override func encode(_ key: String, value: Any?) -> Any? {
        switch key {
        case "sharedWith":
            return self.sharedWith?.firestoreValue
        default:
            break
        }
        return nil
    }

    override func decode(_ key: String, value: Any?) -> Bool {
        switch key {
        case "sharedWith":
            self.sharedWith = Audience(firestoreValue: value)
        default:
            break
        }
        return false
    }
}

@objcMembers class Message: Pring.Object {
    dynamic var content: MessageContent?
    dynamic var from: Pring.Reference<Friend> = .init()

    override func encode(_ key: String, value: Any?) -> Any? {
        switch key {
        case "content":
            return self.content?.firestoreValue
        default:
            break
        }
        return nil
    }

    override func decode(_ key: String, value: Any?) -> Bool {
        switch key {
        case "content":
            self.content = MessageContent(firestoreValue: value)
        default:
            break
        }
        return false
    }
}

@objcMembers class Attachment: Pring.Object {
    dynamic var title: String?
    dynamic var content: AttachmentContent?

    override func encode(_ key: String, value: Any?) -> Any? {
        switch key {
        case "content":
            return self.content?.firestoreValue
        default:
            break
        }
        return nil
    }

    override func decode(_ key: String, value: Any?) -> Bool {
        switch key {
        case "content":
            self.content = AttachmentContent(firestoreValue: value)
        default:
            break
        }
        return false
    }
}

@objcMembers class Friend: Pring.Object {
    dynamic var username: String?
    dynamic var displayName: String?
    dynamic var avatar: Avatar?
    dynamic var friendsSinice: Date?

    override func encode(_ key: String, value: Any?) -> Any? {
        switch key {
        case "avatar":
            return self.avatar?.rawValue
        default:
            break
        }
        return nil
    }

    override func decode(_ key: String, value: Any?) -> Bool {
        switch key {
        case "avatar":
          if let value = value as? [String: Any] {
            self.avatar = Avatar(id: "\(0)", value: value)
            return true
          }
        default:
            break
        }
        return false
    }
}
