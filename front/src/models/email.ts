/**
 * Represents an email message.
 */
export interface Email {
  /**
   * Unique identifier for the email.
   */
  id: string;

  /**
   * Identifier for the message, typically used for threading.
   */
  message_id: string;

  /**
   * Email address of the sender.
   */
  from: string;

  /**
   * Email address of the recipient.
   */
  to: string;

  /**
   * Subject line of the email.
   */
  subject: string;

  /**
   * Date the email was sent.
   */
  date: string;

  /**
   * Body content of the email. Optional.
   */
  body?: string;
}
