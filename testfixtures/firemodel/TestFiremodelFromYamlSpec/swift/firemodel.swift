// DO NOT EDIT - Code generated by firemodel (dev).

import Foundation
import Pring

// TODO: Add documentation to Direction.
@objc enum Direction: Int {
    // TODO: Add documentation to left.
    case left
    // TODO: Add documentation to right.
    case right
    // TODO: Add documentation to up.
    case up
    // TODO: Add documentation to down.
    case down
}

extension Direction: CustomDebugStringConvertible {
    init?(firestoreValue value: Any?) {
        guard let value = value as? String else {
            return nil
        }
        switch value {
        case "LEFT":
            self = .left
        case "RIGHT":
            self = .right
        case "UP":
            self = .up
        case "DOWN":
            self = .down
        default:
            return nil
        }
    }

    var firestoreValue: String? {
        switch self {
        case .left:
            return "LEFT"
        case .right:
            return "RIGHT"
        case .up:
            return "UP"
        case .down:
            return "DOWN"
        default:
            return nil
        }
    }

    var debugDescription: String { return firestoreValue ?? "<INVALID>" }
}

// A Test is a test model.
@objcMembers class TestModel: Pring.Object {
    static var userId: String = ""
    override class var path: String { return "users/\(userId)/test_models" }

    // The name.
    dynamic var name: String?
    // The age.
    dynamic var age: Int = 0
    // The number pi.
    dynamic var pi: Float = 0.0
    // The birth date.
    dynamic var birthdate: Date = Date()
    // True if it is good.
    dynamic var isGood: Bool = false
    // TODO: Add documentation to data.
    dynamic var data: Data?
    // TODO: Add documentation to friend.
    dynamic var friend: Pring.Reference<TestModel> = .init()
    // TODO: Add documentation to location.
    dynamic var location: Pring.GeoPoint?
    // TODO: Add documentation to colors.
    dynamic var colors: [String] = []
    // TODO: Add documentation to meta.
    dynamic var meta: [AnyHashable: Any] = [:]
    // TODO: Add documentation to direction.
    dynamic var direction: Direction?
    // TODO: Add documentation to testFile.
    dynamic var testFile: Pring.File?
    // TODO: Add documentation to url.
    dynamic var url: URL?
    // TODO: Add documentation to nested.
    dynamic var nested: Pring.NestedCollection<TestModel> = []

    
}
